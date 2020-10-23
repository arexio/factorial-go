# Factorial API in Go

The aim of this project is to create a full SDK in Go for using the Factorial API. In order to know which services you can consume please refer to the factorial API documentation https://docs.factorialhr.com/reference.

## Getting Started

We recomend to follow first of all the getting started guide from factorial https://docs.factorialhr.com/docs. After following this guide you will have created a new OAuth application with the information need to start using this SDK. The next fields are the ones you need in order to start using factorial.

```
CLIENT_ID="---- Your client ID ----"
CLIENT_SECRET="--- Your client secret ---"
SCOPES="read,write"
REDIRECT_URL="--- Your redirect url, sample (https://7cad0b374498.ngrok.io/auth/factorial/callback) ----"
```

We provide a Makefile that will allow you to start our different examples, we have three examples:

* example/employee: this is a short example of how to use our factorial client for retrieve a list of employees
* example/server: this example will setup and run a server from where you can see how we should handle the OAuth2 process, as well you will find there all the endpoints and you can try them.
* example/persistence: this example will setup and run a server, in the same we had for the server example, but on this case we use a token repository for persist and refresh our token.

## How to use

On the next snippet you can see how we build our client

```
    // Build the Oauth provider
    provider = factorial.NewOAuthProvider(
		factorial.WithClientID(clientID),
		factorial.WithClientSecret(clientSecret),
		factorial.WithScopes(scopes),
		factorial.WithRedirectURL(redirectURL),
	)

    // Build the factorial client
    cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.Client(token)),
	)
	if err != nil {
		// Track error
	}

    // Use the factorial client for retrieve a list of employees
    employees, err := cl.ListEmployees()
	if err != nil {
		// Track error
	}
```

The next snippet will show you how to use the factorial client based on our token repository

```
    // Build the Oauth provider
    provider = factorial.NewOAuthProvider(
		factorial.WithClientID(clientID),
		factorial.WithClientSecret(clientSecret),
		factorial.WithScopes(scopes),
		factorial.WithRedirectURL(redirectURL),
	)

    // You can find the token repository on the repository.go file
    // as well a sample implementation on the persistence sample
	repo = NewMemoryRepository()

    // Build a new OAuth client with a custom Source with repository
    cl, err := factorial.New(
		factorial.WithOAuth2Client(provider.ClientWithSource(
			factorial.NewTokenSource(
				repo,
				uuid.Nil,
				provider,
			),
		)),
	)
	if err != nil {
		// Track error
	}

    // Use the factorial client for retrieve a list of employees
    employees, err := cl.ListEmployees()
	if err != nil {
		// Track error
	}
```