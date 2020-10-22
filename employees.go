package factorial

import (
	"encoding/json"
)

const (
	employeeURL = "/api/v1/employees"
)

// Employee contains all the employee information.
type Employee struct {
	ID                   int    `json:"id"`
	BirthdayOn           string `json:"birthday_on"`
	StartDate            string `json:"start_date"`
	Email                string `json:"email"`
	FullName             string `json:"full_name"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	ManagerID            int    `json:"manager_id"`
	Role                 string `json:"role"`
	TimeoffManagerID     int    `json:"timeoff_manager_id"`
	TerminatedOn         string `json:"terminated_on"`
	PhoneNumber          string `json:"phone_number"`
	Gender               string `json:"gender"`
	Nationality          string `json:"nationality"`
	BankNumber           string `json:"bank_number"`
	Country              string `json:"country"`
	City                 string `json:"city"`
	State                string `json:"state"`
	PostalCode           string `json:"postal_code"`
	AddresLine1          string `json:"address_line_1"`
	AddressLine2         string `json:"address_line_2"`
	SocialSecurityNumber string `json:"social_security_number"`
	CompanyHolidayIDs    []int  `json:"company_holiday_ids"`
	Identifier           string `json:"identifier"`      // National identification number
	IdentifierType       string `json:"identifier_type"` //Type of national identification. Possible value: DNI, NIE, Passport
	Hiring               Hiring `json:"hiring"`
	LocationID           int    `json:"location_id"`
	TeamIDs              []int  `json:"team_ids"`
}

// Hiring encapsulates the hiring details of an employee.
type Hiring struct {
	BaseCompensationAmountInCents int    `json:"base_compensation_amount_in_cents"`
	BaseCompensationType          string `json:"base_compensation_type"` // Compensation recurrence. Possible values: hourly, monthly, yearly.
}

// CreateEmployeeRequest is the object for create an employee.
type CreateEmployeeRequest struct {
	BirthdayOn       string `json:"birthday_on,omitempty"`
	StartDate        string `json:"start_date,omitempty"`
	Email            string `json:"email"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	ManagerID        int    `json:"manager_id,omitempty"`
	Role             string `json:"role,omitempty"`
	TimeoffManagerID int    `json:"timeoff_manager_id,omitempty"`
	TerminatedOn     string `json:"terminated_on,omitempty"`
	TerminatedReason string `json:"terminated_reason,omitempty"`
}

// CreateEmployee creates a new Employee in your company.
// Restricted to admin users.
func (c Client) CreateEmployee(e CreateEmployeeRequest) (Employee, error) {
	var employee Employee

	bytes, err := json.Marshal(e)
	if err != nil {
		return employee, err
	}

	resp, err := c.post(employeeURL, bytes)
	if err != nil {
		return employee, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&employee); err != nil {
		return employee, err
	}

	return employee, nil
}

// GetEmployee gets all information for an employee.
func (c Client) GetEmployee(id string) (Employee, error) {
	var employee Employee

	resp, err := c.get(employeeURL+"/"+id, nil)
	if err != nil {
		return employee, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&employee); err != nil {
		return employee, err
	}

	return employee, nil
}

// ListEmployees gets all employees from your company.
// Only admins can see all the employees' information, regular users will get a restricted version of the payload as a response.
func (c Client) ListEmployees() ([]Employee, error) {
	var employees []Employee

	resp, err := c.get(employeeURL, nil)
	if err != nil {
		return employees, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&employees); err != nil {
		return employees, err
	}

	return employees, nil
}

// TerminateEmployee terminates an existing Employee.
// This is not a hard delete but simply a flag toggle in the Employee model.
// Restricted to admin users.
func (c Client) TerminateEmployee(id, date, reason string) (Employee, error) {
	var employee Employee

	bytes, err := json.Marshal(map[string]string{
		"terminated_on":     date,
		"terminated_reason": reason,
	})
	if err != nil {
		return employee, err
	}

	resp, err := c.post(employeeURL+"/"+id+"/terminate", bytes)
	if err != nil {
		return employee, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&employee); err != nil {
		return employee, err
	}

	return employee, nil
}

// UpdateEmployeeRequest is the object for update an employee.
type UpdateEmployeeRequest struct {
	BirthdayOn       string `json:"birthday_on,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	ManagerID        int    `json:"manager_id,omitempty"`
	Role             string `json:"role,omitempty"`
	TimeoffManagerID int    `json:"timeoff_manager_id,omitempty"`
}

// UpdateEmployee updates an existing Employee.
// Admin users can update all parameters whereas regular users can only update a subset of attributes.
func (c Client) UpdateEmployee(id string, e UpdateEmployeeRequest) (Employee, error) {
	var employee Employee

	bytes, err := json.Marshal(e)
	if err != nil {
		return employee, err
	}

	resp, err := c.put(employeeURL+"/"+id, bytes)
	if err != nil {
		return employee, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&employee); err != nil {
		return employee, err
	}

	return employee, nil
}

// UnterminateEmployee removes the termination date of an Employee.
// Restricted to admin users.
func (c Client) UnterminateEmployee(id string) (Employee, error) {
	var employee Employee

	resp, err := c.post(employeeURL+"/"+id+"/terminate", nil)
	if err != nil {
		return employee, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&employee); err != nil {
		return employee, err
	}

	return employee, nil
}
