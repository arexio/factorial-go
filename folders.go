package factorial

import (
	"encoding/json"
	"net/url"
)

const (
	folderURL = "/api/v1/folders"
)

// Folder contains all the folder information
type Folder struct {
	ID        int    `json:"id"`
	CompanyID int    `json:"company_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateFolderRequest keeps the information needed
// for create a new folder
type CreateFolderRequest struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// UpdateFolderRequest keeps the information needed
// for update a folder
type UpdateFolderRequest struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// CreateFolder creates a new folder in your company
func (c Client) CreateFolder(f CreateFolderRequest) (Folder, error) {
	var folder Folder

	bytes, err := json.Marshal(f)
	if err != nil {
		return folder, err
	}

	resp, err := c.post(folderURL, bytes)
	if err != nil {
		return folder, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&folder); err != nil {
		return folder, err
	}

	return folder, nil
}

// GetFolder gets all information for the given folderID
func (c Client) GetFolder(id string) (Folder, error) {
	var folder Folder

	resp, err := c.get(folderURL+"/"+id, nil)
	if err != nil {
		return folder, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&folder); err != nil {
		return folder, err
	}

	return folder, nil
}

// ListFolders gets all the folder from you company
// you can filter this list by name and active
func (c Client) ListFolders(filter url.Values) ([]Folder, error) {
	var folders []Folder

	resp, err := c.get(folderURL, filter)
	if err != nil {
		return folders, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&folders); err != nil {
		return folders, err
	}

	return folders, nil
}

// UpdateFolder update the given folder id with the given
// request data
func (c Client) UpdateFolder(id string, f UpdateFolderRequest) (Folder, error) {
	var folder Folder

	bytes, err := json.Marshal(f)
	if err != nil {
		return folder, err
	}

	resp, err := c.put(folderURL+"/"+id, bytes)
	if err != nil {
		return folder, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&folder); err != nil {
		return folder, err
	}

	return folder, nil
}
