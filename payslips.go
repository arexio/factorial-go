package factorial

import (
	"encoding/json"
	"net/url"
)

const (
	payslipURL = "/api/v1/payslips"
)

// Payslip keeps the basic information related
// with payslips in Factorial
type Payslip struct {
	ID                    int    `json:"id"`
	BaseCotizationInCents int    `json:"base_cotization_in_cents"`
	BaseIRPFInCents       int    `json:"base_irpf_in_cents"`
	GrossSalaryInCents    string `json:"gross_salary_in_cents"`
	NetSalaryInCents      string `json:"net_salary_in_cents"`
	IRPFInCents           bool   `json:"irpf_in_cents"`
	IRPFPercentage        string `json:"irpf_percentage"`
	IsLastPayslip         bool   `json:"is_last_payslip"`
	StartDate             string `json:"start_date"`
	EndDate               string `json:"end_date"`
	EmployeeID            int    `json:"employee_id"`
	Status                string `json:"status"`
}

// ListPayslips gets all the payslips from your company
// you can filter this list by status, specific year and month,
// or get all payslips from a specific month and year
// with the from param: "from: {month: 12, year: 2019}".
func (c Client) ListPayslips(filter url.Values) ([]Payslip, error) {
	var payslips []Payslip

	resp, err := c.get(payslipURL, filter)
	if err != nil {
		return payslips, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&payslips); err != nil {
		return payslips, err
	}

	return payslips, nil
}
