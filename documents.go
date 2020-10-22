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

// CreateDocumentRequest will hold the basic information
// needed for create a new document in Factorial
type CreateDocumentRequest struct {
	Public            bool   `json:"public"`
	EmployeeID        int    `json:"employee_id"`
	File              string `json:"file"`
	FileName          string `json:"filename"`
	FolderID          int    `json:"folder_id"`
	RequestESignature bool   `json:"request_esignature"`
	Signees           []int  `json:"signees"`
}

// UpdateDocumentRequest will hold the basic information
// for update a given document
type UpdateDocumentRequest struct {
	Public            bool  `json:"public"`
	EmployeeID        int   `json:"employee_id"`
	FolderID          int   `json:"folder_id"`
	RequestESignature bool  `json:"request_esignature"`
	Signees           []int `json:"signees"`
}

// CreateDocument creates a new document in Factorial
func (c Client) CreateDocument(d CreateDocumentRequest) (Document, error) {
	var document Document

	bytes, err := json.Marshal(d)
	if err != nil {
		return document, err
	}

	resp, err := c.post(documentURL, bytes)
	if err != nil {
		return document, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&document); err != nil {
		return document, err
	}

	return document, nil
}

// DeleteDocument will delete the given documentID
func (c Client) DeleteDocument(id string) error {
	_, err := c.delete(documentURL + "/" + id)
	if err != nil {
		return err
	}

	return nil
}

// GetDocument return the document saved in Factorial with
// the given id
func (c Client) GetDocument(id string) (Document, error) {
	var document Document

	resp, err := c.get(documentURL+"/"+id, nil)
	if err != nil {
		return document, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&document); err != nil {
		return document, err
	}

	return document, nil
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

// UpdateDocument update the given document id with the given data
func (c Client) UpdateDocument(id string, d UpdateDocumentRequest) (Document, error) {
	var document Document

	bytes, err := json.Marshal(d)
	if err != nil {
		return document, err
	}

	resp, err := c.put(documentURL+"/"+id, bytes)
	if err != nil {
		return document, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&document); err != nil {
		return document, err
	}

	return document, nil
}
