package main

import (
	"github.com/arexio/factorial-go"
	"github.com/gofrs/uuid"
	"golang.org/x/oauth2"
)

type memoryRepository struct {
	tokens map[uuid.UUID]*oauth2.Token
}

// NewMemoryRepository will build a new factorial.TokenRepository
// based on a in memory implementation
func NewMemoryRepository() factorial.TokenRepository {
	return &memoryRepository{
		tokens: make(map[uuid.UUID]*oauth2.Token),
	}
}

func (m *memoryRepository) SaveToken(id uuid.UUID, t *oauth2.Token) error {
	m.tokens[id] = t
	return nil
}

func (m *memoryRepository) UpdateToken(id uuid.UUID, t *oauth2.Token) error {
	m.tokens[id] = t
	return nil
}

func (m *memoryRepository) GetToken(id uuid.UUID) (*oauth2.Token, error) {
	return m.tokens[id], nil
}
