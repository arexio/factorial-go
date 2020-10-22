package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"

	"github.com/arexio/factorial-go"
)

var (
	clientID, clientSecret, redirectURL string
	scopes                              []string
	provider                            *factorial.OAuthProvider
	token                               *oauth2.Token
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panicln("No .env file found")
	}

	clientID = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	scopes = strings.Split(os.Getenv("SCOPES"), ",")
	redirectURL = os.Getenv("REDIRECT_URL")
	provider = factorial.NewOAuthProvider(
		factorial.WithClientID(clientID),
		factorial.WithClientSecret(clientSecret),
		factorial.WithScopes(scopes),
		factorial.WithRedirectURL(redirectURL),
	)
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/auth/factorial", StartFactorialOAuthHandler)
	r.HandleFunc("/auth/factorial/callback", FactorialOAuthCallbackHandler)
	r.HandleFunc("/employees", EmployeesHandler)
	r.HandleFunc("/folders", FoldersHandler)
	r.HandleFunc("/leave_types", LeaveTypesHandler)
	r.HandleFunc("/leaves", LeavesHandler)
	r.HandleFunc("/webhooks", WebhooksHandler)

	staticRouter := r.PathPrefix("/static/")
	staticRouter.Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./public"))))

	http.Handle("/", r)

	srv := &http.Server{
		Addr: "0.0.0.0:3000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	log.Println("Server running on port 3000")

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

// HomeHandler will be the base handler in where we will show information about
// token and different actions you can do
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	var err error

	if token == nil {
		t, err = template.New("index").Parse(indexTemplate)
	} else {
		t, err = template.New("connected").Parse(connectedTemplate)
	}
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, nil)
}

// StartFactorialOAuthHandler is the handler that will start the process of Auth with
// the Factorial platform
func StartFactorialOAuthHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, provider.GetAuthURL("uniq_state"), http.StatusFound)
}

// FactorialOAuthCallbackHandler is the handler in where we are going to receive a
// successful callback with a code that can we use to get our user token
func FactorialOAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	token, err = provider.GetTokenFromCode(r.FormValue("code"))
	if err != nil {
		log.Panic(err)
	}
	t, err := template.New("connected").Parse(connectedTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, token)
}

// EmployeesHandler is the handler used for get all the employees
// and print them on a list template
func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	employees, err := cl.ListEmployees()
	if err != nil {
		log.Panicln("Error while getting employees", err)
	}
	t, err := template.New("employees").Parse(employeesTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Employees []factorial.Employee
	}{
		Employees: employees,
	})
}

// FoldersHandler is the handler used for get all the folders
// and print them on a list template
func FoldersHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	folders, err := cl.ListFolders(nil)
	if err != nil {
		log.Panicln("Error while getting folders", err)
	}
	t, err := template.New("folders").Parse(foldersTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Folders []factorial.Folder
	}{
		Folders: folders,
	})
}

// LeaveTypesHandler is the handler used for get all the leave types
// and print them on a list template
func LeaveTypesHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	leaveTypes, err := cl.ListLeaveTypes()
	if err != nil {
		log.Panicln("Error while getting leave types", err)
	}
	t, err := template.New("leaveTypes").Parse(leaveTypesTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		LeaveTypes []factorial.LeaveType
	}{
		LeaveTypes: leaveTypes,
	})
}

// LeavesHandler is the handler used for get all the leaves
// and print them on a list template
func LeavesHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	leaves, err := cl.ListLeaves()
	if err != nil {
		log.Panicln("Error while getting leaves", err)
	}
	t, err := template.New("leaves").Parse(leavesTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Leaves []factorial.Leave
	}{
		Leaves: leaves,
	})
}

// WebhooksHandler is the handler used for get all the webhooks
// and print them on a list template
func WebhooksHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	webhooks, err := cl.ListWebhooks()
	if err != nil {
		log.Panicln("Error while getting webhooks", err)
	}
	t, err := template.New("webhooks").Parse(webhooksTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Webhooks []factorial.Webhook
	}{
		Webhooks: webhooks,
	})
}
