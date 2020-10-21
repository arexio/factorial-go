package main

import (
	"fmt"
	"log"

	"golang.org/x/oauth2"

	factorial "github.com/arexio/factorial-go"
)

func main() {
	// Fill the empty values with your provided
	// factorial app values
	provider := factorial.NewOAuthProvider(
		factorial.WithClientID(""),
		factorial.WithClientSecret(""),
		factorial.WithScopes([]string{}),
		factorial.WithRedirectURL(""),
	)

	// Fill the empty token with the one provided
	// from factorial
	token := &oauth2.Token{}

	cli, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		log.Println(err)
		return
	}

	employees, err := cli.ListEmployees()
	if err != nil {
		log.Println(err)
		return
	}

	for _, e := range employees {
		fmt.Printf("%+v\n", e)
	}
}
