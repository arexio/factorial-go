package factorial

import (
	"github.com/gofrs/uuid"
	"golang.org/x/oauth2"
)

// TokenRepository interface will be used for persist
// our oauth token
type TokenRepository interface {
	SaveToken(id uuid.UUID, t *oauth2.Token) error
	UpdateToken(id uuid.UUID, t *oauth2.Token) error
	GetToken(id uuid.UUID) (*oauth2.Token, error)
}

// tokenRefresher implements oauth2.TokenSource
// with a custom repository
type tokenRefresher struct {
	repo     TokenRepository
	token    *oauth2.Token
	id       uuid.UUID
	provider *OAuthProvider
}

// NewTokenSource will build a new token source with the given criteria
func NewTokenSource(repo TokenRepository, token *oauth2.Token, id uuid.UUID, provider *OAuthProvider) oauth2.TokenSource {
	return &tokenRefresher{
		repo:     repo,
		token:    token,
		id:       id,
		provider: provider,
	}
}

// Token method is the custom implementation of the refresh token process using
// a token repo as a base
func (t *tokenRefresher) Token() (*oauth2.Token, error) {
	if !t.token.Valid() {
		token, err := t.provider.RefreshToken(t.token)
		if err != nil {
			return nil, err
		}
		if err = t.repo.UpdateToken(t.id, token); err != nil {
			return nil, err
		}
		return token, nil
	}
	return t.token, nil
}
