package factorial

import "encoding/json"

const (
	locationURL = "/api/v1/locations"
)

// Location keeps the basic information related with a location
type Location struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Country            string `json:"country"` // example: es
	PhoneNumber        string `json:"phone_number"`
	State              string `json:"state"` // State/province code. E.g. 'ct'
	City               string `json:"city"`  // City name. E.g. 'barcelona'
	AddressLine1       string `json:"address_line_1"`
	AddressLine2       string `json:"address_line_2"`
	PostalCode         string `json:"postal_code"`
	CompanyHolidaysIDs []int  `json:"company_holidays_ids"`
}

// GetLocation will get the location linked to the given id
func (c Client) GetLocation(id string) (Location, error) {
	var location Location

	resp, err := c.get(locationURL+"/"+id, nil)
	if err != nil {
		return location, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return location, err
	}

	return location, nil
}

// ListLocations will get all the location saved into Factorial
func (c Client) ListLocations() ([]Location, error) {
	var locations []Location

	resp, err := c.get(locationURL, nil)
	if err != nil {
		return locations, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		return locations, err
	}

	return locations, nil
}
