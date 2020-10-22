package factorial

import (
	"encoding/json"
)

const (
	companyHolidayURL = "/api/v1/company_holidays"
)

// CompanyHoliday holds the basic information related
// with the company holidays in Factorial
type CompanyHoliday struct {
	ID          int    `json:"id"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Date        string `json:"date"`
	HalfDay     string `json:"half_day"`
	LocationID  int    `json:"location_id"`
}

// GetCompanyHoliday will get the company holiday linked
// to the given id
func (c Client) GetCompanyHoliday(id string) (CompanyHoliday, error) {
	var companyHoliday CompanyHoliday

	resp, err := c.get(companyHolidayURL+"/"+id, nil)
	if err != nil {
		return companyHoliday, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&companyHoliday); err != nil {
		return companyHoliday, err
	}

	return companyHoliday, nil
}

// ListCompanyHolidays will get all the company holidays saved
// in Factorial
func (c Client) ListCompanyHolidays() ([]CompanyHoliday, error) {
	var companyHolidays []CompanyHoliday

	resp, err := c.get(companyHolidayURL, nil)
	if err != nil {
		return companyHolidays, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&companyHolidays); err != nil {
		return companyHolidays, err
	}

	return companyHolidays, nil
}
