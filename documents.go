package factorial

import (
	"encoding/json"
	"net/url"
)

const (
	documentURL = "/api/v1/documents"
)

// Document keeps the basic information related
// with documents in Factorial
type Document struct {
	ID         int    `json:"id"`
	EmployeeID int    `json:"employee_id"`
	CompanyID  int    `json:"company_id"`
	FolderID   int    `json:"folder_id"`
	File       string `json:"file"`
	FileName   string `json:"filename"`
	Public     bool   `json:"public"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

// ListDocuments gets all the documents from your company
// you can filter this list by employee_id and folder_id
func (c Client) ListDocuments(filter url.Values) ([]Document, error) {
	var documents []Document

	resp, err := c.get(documentURL, filter)
	if err != nil {
		return documents, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&documents); err != nil {
		return documents, err
	}

	return documents, nil
}
