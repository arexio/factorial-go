package factorial

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	// GrantTypeGetToken is the grant_type value when getting the acces token.
	GrantTypeGetToken = "authorization_code"
	// GrantTypeRefreshToken is the grant_type value when refreshing access token.
	GrantTypeRefreshToken = "refresh_token"
)

// OAuth2GetToken is the request object for get the access token.
type OAuth2GetToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	GrantType    string `json:"grant_type"`
	RedirectURI  string `json:"redirect_uri"`
}

// OAuth2RefreshToken is the request object for refresh the access token.
type OAuth2RefreshToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

// OAuth2TokenResponse is the response object for get/refresh the access token.
type OAuth2TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

// GetToken returns the access token.
func GetToken(apiURL string, body OAuth2GetToken) (OAuth2TokenResponse, error) {
	var data OAuth2TokenResponse

	buff := new(bytes.Buffer)
	if err := json.NewEncoder(buff).Encode(body); err != nil {
		return data, err
	}

	return doRequest(apiURL, buff)
}

// RefreshToken returns the "refreshed" access token.
func RefreshToken(apiURL string, body OAuth2RefreshToken) (OAuth2TokenResponse, error) {
	var data OAuth2TokenResponse

	buff := new(bytes.Buffer)
	if err := json.NewEncoder(buff).Encode(body); err != nil {
		return data, err
	}

	return doRequest(apiURL, buff)
}

func doRequest(apiURL string, buff *bytes.Buffer) (OAuth2TokenResponse, error) {
	var data OAuth2TokenResponse

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%soauth/token", apiURL), buff)
	if err != nil {
		return data, err

	}
	req.Header.Set("Content-Type", "application/json")

	c := http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := c.Do(req)
	if err != nil {
		return data, err
	}
	if resp.StatusCode != http.StatusOK {
		return data, fmt.Errorf("error getting token: %s", resp.Status)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)

	return data, err
}
