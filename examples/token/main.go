package main

import (
	"log"

	factorial "github.com/charly3pins/factorial-go"
)

const (
	apiURL            = "https://api.factorialhr.com/"
	clientID          = "-"
	clientSecret      = "-"
	redirectURI       = "-"
	authorizationCode = "-"
)

func main() {
	req := factorial.OAuth2GetToken{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         authorizationCode,
		GrantType:    factorial.GrantTypeGetToken,
		RedirectURI:  redirectURI,
	}

	resp, err := factorial.GetToken(apiURL, req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v", resp)
}
