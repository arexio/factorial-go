package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
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
	r.HandleFunc("/documents", DocumentsHandler)
	r.HandleFunc("/hiring_versions", HiringVersionsHandler)
	r.HandleFunc("/company_holidays", CompanyHolidaysHandler)
	r.HandleFunc("/payslips", PayslipsHandler)
	r.HandleFunc("/locations", LocationsHandler)
	r.HandleFunc("/teams", TeamsHandler)
	r.HandleFunc("/shifts", ShiftsHandler)
	r.HandleFunc("/clocks", ClocksHandler)
	r.HandleFunc("/clock_in", ClockInHandler)
	r.HandleFunc("/clock_out", ClockOutHandler)

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
	t.Execute(w, struct {
		Webhooks []factorial.Webhook
	}{
		Webhooks: webhooks,
	})
}

// DocumentsHandler is the handler used for get all the documents
// and print them on a list template
func DocumentsHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	documents, err := cl.ListDocuments(nil)
	if err != nil {
		log.Panicln("Error while getting documents", err)
	}
	t, err := template.New("documents").Parse(documentsTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Documents []factorial.Document
	}{
		Documents: documents,
	})
}

// HiringVersionsHandler is the handler used for get all the hiring versions
// and print them on a list template
func HiringVersionsHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	hiringVersions, err := cl.ListHiringVersions(nil)
	if err != nil {
		log.Panicln("Error while getting hiring versions", err)
	}
	log.Println("[DEBUG] hiring versions", len(hiringVersions))
	t, err := template.New("hiringVersions").Parse(hiringVersionsTemplate)
	t.Execute(w, struct {
		HiringVersions []factorial.HiringVersion
	}{
		HiringVersions: hiringVersions,
	})
}

// PayslipsHandler is the handler used for get all the payslips
// and print them on a list template
func PayslipsHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	payslips, err := cl.ListPayslips(nil)
	if err != nil {
		log.Panicln("Error while getting payslips", err)
	}
	log.Println("[DEBUG] payslips", len(payslips))
	t, err := template.New("payslips").Parse(payslipsTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Payslips []factorial.Payslip
	}{
		Payslips: payslips,
	})
}

// CompanyHolidaysHandler is the handler for get all the company holidays
// and print them on a list template
func CompanyHolidaysHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	companyHolidays, err := cl.ListCompanyHolidays()
	if err != nil {
		log.Panicln("Error while getting company holidays", err)
	}
	t, err := template.New("companyHolidays").Parse(companyHolidaysTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		CompanyHolidays []factorial.CompanyHoliday
	}{
		CompanyHolidays: companyHolidays,
	})
}

// LocationsHandler is the handler used for get all the locations
// and print them on a list template
func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	locations, err := cl.ListLocations()
	if err != nil {
		log.Panicln("Error while getting locations", err)
	}
	log.Println("[DEBUG] locations", len(locations))
	t, err := template.New("locations").Parse(locationsTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Locations []factorial.Location
	}{
		Locations: locations,
	})
}

// TeamsHandler is the handler used for get all the teams
// and print them on a list template
func TeamsHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	teams, err := cl.ListTeams()
	if err != nil {
		log.Panicln("Error while getting teams", err)
	}
	log.Println("[DEBUG] teams", len(teams))
	t, err := template.New("teams").Parse(teamsTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Teams []factorial.Team
	}{
		Teams: teams,
	})
}

// ShiftsHandler is the handler used for get all the shifts
// and print them on a list template
func ShiftsHandler(w http.ResponseWriter, r *http.Request) {
	cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Panic(err)
	}

	shifts, err := cl.ListShifts(nil)
	if err != nil {
		log.Panicln("Error while getting shifts", err)
	}
	log.Println("[DEBUG] shifts", len(shifts))
	t, err := template.New("shifts").Parse(shiftsTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Shifts []factorial.Shift
	}{
		Shifts: shifts,
	})
}

// ClocksHandler is the handler used for test the clock_in and clock_out
// and print the result
func ClocksHandler(w http.ResponseWriter, r *http.Request) {
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
	log.Println("[DEBUG] employees", len(employees))
	t, err := template.New("clocks").Parse(clocksTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Employees []factorial.Employee
	}{
		Employees: employees,
	})
}

// ClockInHandler is the handler used for clock_in to Factorial
// and print the result
func ClockInHandler(w http.ResponseWriter, r *http.Request) {
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
	log.Println("[DEBUG] employees", len(employees))

	r.ParseForm()

	employeeID, err := strconv.Atoi(r.PostForm["employees"][0])
	if err != nil {
		log.Panicln("Error while converting employeeID to int", err)
	}

	req := factorial.ClockInRequest{
		EmployeeID: employeeID,
		Now:        time.Now().Format(time.RFC3339),
	}
	shift, err := cl.ClockIn(req)
	if err != nil {
		log.Panicln("Error while clocking in", err)
	}
	log.Println("[DEBUG] clock_in shift", shift)
	t, err := template.New("clocks").Parse(clocksTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Employees     []factorial.Employee
		ClockIntShift factorial.Shift
	}{
		Employees:     employees,
		ClockIntShift: shift,
	})
}

// ClockOutHandler is the handler used for clock_out to Factorial
// and print the result
func ClockOutHandler(w http.ResponseWriter, r *http.Request) {
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
	log.Println("[DEBUG] employees", len(employees))

	r.ParseForm()

	employeeID, err := strconv.Atoi(r.PostForm["employees"][0])
	if err != nil {
		log.Panicln("Error while converting employeeID to int", err)
	}

	req := factorial.ClockOutRequest{
		EmployeeID: employeeID,
		Now:        time.Now().Format(time.RFC3339),
	}
	shift, err := cl.ClockOut(req)
	if err != nil {
		log.Panicln("Error while clocking out", err)
	}
	log.Println("[DEBUG] clock_out shift", shift)
	t, err := template.New("clocks").Parse(clocksTemplate)
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, struct {
		Employees     []factorial.Employee
		ClockOutShift factorial.Shift
	}{
		Employees:     employees,
		ClockOutShift: shift,
	})
}
