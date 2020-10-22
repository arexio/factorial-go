package factorial

import (
	"encoding/json"
	"net/url"
)

const (
	hiringVersionURL = "/api/v1/hiring_versions"
)

// HiringVersion keeps the basic information related
// with hiring versions in Factorial
type HiringVersion struct {
	ID                            int    `json:"id"`
	EffectiveOn                   string `json:"effective_on"` // Date from which this contract is valid
	EmployeeID                    int    `json:"employee_id"`
	BaseCompensationAmountInCents int    `json:"base_compensation_amount_in_cents"` // Gross salary in cents
	BaseCompensationType          string `json:"base_compensation_type"`            // Gross salary recurrence type
	StartDate                     string `json:"start_date"`                        // Employee starting date
	EndDate                       string `json:"end_date"`                          // Employee end date
	JobTitle                      string `json:"job_title"`
	WorkingHoursInCents           int    `json:"working_hours_in_cents"`
	WorkingPeriodUnit             string `json:"working_period_unit"` //Working hours recurrence type
}

// ListHiringVersions gets all the hiring versions from employees
// you can filter this list by employee_id
func (c Client) ListHiringVersions(filter url.Values) ([]HiringVersion, error) {
	var hiringVersions []HiringVersion

	resp, err := c.get(hiringVersionURL, filter)
	if err != nil {
		return hiringVersions, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&hiringVersions); err != nil {
		return hiringVersions, err
	}

	return hiringVersions, nil
}
