package factorial

import (
	"encoding/json"
	"net/url"
)

const (
	shiftURL    = "/api/v1/shifts"
	clockInURL  = shiftURL + "/clock_in"
	clockOutURL = shiftURL + "/clock_out"
)

// Shift keeps the basic information related
// with shifts in Factorial
type Shift struct {
	ID           int    `json:"id"`
	Day          int    `json:"day"`
	Month        int    `json:"month"`
	Year         int    `json:"year"`
	ClockIn      string `json:"clock_in"`
	ClockOut     string `json:"clock_out"`
	EmployeeID   int    `json:"employee_id"`
	Observations string `json:"observations"`
}

// ClockInRequest will hold the basic information
// needed for create a new shift (ClockIn) in Factorial
type ClockInRequest struct {
	Now        string `json:"now"`
	EmployeeID int    `json:"employee_id"`
}

// ClockOutRequest will hold the basic information
// needed for create a new shift (ClockOut) in Factorial
type ClockOutRequest struct {
	Now        string `json:"now"`
	EmployeeID int    `json:"employee_id"`
}

// UpdateShiftRequest will hold the basic information
// for update a given shift.
// Restricted to the user's own shifts.
type UpdateShiftRequest struct {
	ClockIn      string `json:"clock_in"`
	ClockOut     string `json:"clock_out"`
	Observations string `json:"observations"`
}

// ClockIn creates a new Shift with the provided time and for the requested employee.
// If an open Shift already exists, this endpoint will return an error with code 422.
func (c Client) ClockIn(cin ClockInRequest) (Shift, error) {
	var shift Shift

	bytes, err := json.Marshal(cin)
	if err != nil {
		return shift, err
	}

	resp, err := c.post(clockInURL, bytes)
	if err != nil {
		return shift, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&shift); err != nil {
		return shift, err
	}

	return shift, nil
}

// ClockOut closes an employee's open shift by setting the clock-out time to the value provided.
// If no open shift exists this endpoint will return an error with status code 422.
func (c Client) ClockOut(cout ClockOutRequest) (Shift, error) {
	var shift Shift

	bytes, err := json.Marshal(cout)
	if err != nil {
		return shift, err
	}

	resp, err := c.post(clockOutURL, bytes)
	if err != nil {
		return shift, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&shift); err != nil {
		return shift, err
	}

	return shift, nil
}

// DeleteShift will delete the given shiftID
func (c Client) DeleteShift(id string) error {
	_, err := c.delete(shiftURL + "/" + id)
	if err != nil {
		return err
	}

	return nil
}

// ListShifts gets all the shifts. Shifts are the unit to control the presence of an employee.
// A Shift has a clock-in and clock-out time (in hours and minutes).
// A shift can be opened by just setting the clock-in time, and later on, closed by updating the clock-out time.
// You can filter this list by year and month.
func (c Client) ListShifts(filter url.Values) ([]Shift, error) {
	var shifts []Shift

	resp, err := c.get(shiftURL, filter)
	if err != nil {
		return shifts, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&shifts); err != nil {
		return shifts, err
	}

	return shifts, nil
}

// UpdateShift update the given shift id with the given data
func (c Client) UpdateShift(id string, d UpdateShiftRequest) (Shift, error) {
	var shift Shift

	bytes, err := json.Marshal(d)
	if err != nil {
		return shift, err
	}

	resp, err := c.put(shiftURL+"/"+id, bytes)
	if err != nil {
		return shift, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&shift); err != nil {
		return shift, err
	}

	return shift, nil
}
