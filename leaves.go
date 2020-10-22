package factorial

import "encoding/json"

const (
	leaveTypeURL = "/api/v1/leave_types"
	leaveURL     = "/api/v1/leaves"
)

// LeaveType contains all the leave type information
type LeaveType struct {
	ID               int    `json:"id"`
	Accrues          bool   `json:"accrues"`           // Whether leaves with this type accrue holidays
	Active           bool   `json:"active"`            // Whether leaves whit this type can be created
	ApprovalRequired bool   `json:"approval_required"` // Whether leaves with this type require approval from timeoff managers
	Attachment       bool   `json:"attachment"`        // Whether leaves with this type accept attachments
	Color            string `json:"color"`             // Identifying color of this leave type
	Identifier       string `json:"identifier"`        // Slug identifying the type of leave type. Only "custom" leave types can be created or modified via the API
	Name             string `json:"name"`
	Visibility       bool   `json:"visibility"` // Whether this leave type is visibile to regular employees
	Workable         bool   `json:"workable"`   // Whether leaves with this type count as working days
}

// CreateLeaveTypeRequest keeps the information needed
// for create a new leave type
type CreateLeaveTypeRequest struct {
	Accrues          bool   `json:"accrues,omitempty"`
	Active           bool   `json:"active,omitempty"`
	ApprovalRequired bool   `json:"approval_required,omitempty"`
	Attachment       bool   `json:"attachment,omitempty"`
	Color            string `json:"color"`
	Name             string `json:"name"`
	Visibility       bool   `json:"visibility,omitempty"`
	Workable         bool   `json:"workable,omitempty"`
}

// UpdateLeaveTypeRequest keeps the information needed
// for update a leave type
type UpdateLeaveTypeRequest struct {
	Accrues          bool   `json:"accrues,omitempty"`
	Active           bool   `json:"active,omitempty"`
	ApprovalRequired bool   `json:"approval_required,omitempty"`
	Attachment       bool   `json:"attachment,omitempty"`
	Color            string `json:"color,omitempty"`
	Name             string `json:"name,omitempty"`
	Visibility       bool   `json:"visibility,omitempty"`
	Workable         bool   `json:"workable,omitempty"`
}

// CreateLeaveType creates a new leave type.
// Restricted to admin users.
func (c Client) CreateLeaveType(lt CreateLeaveTypeRequest) (LeaveType, error) {
	var leaveType LeaveType

	bytes, err := json.Marshal(lt)
	if err != nil {
		return leaveType, err
	}

	resp, err := c.post(leaveTypeURL, bytes)
	if err != nil {
		return leaveType, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&leaveType); err != nil {
		return leaveType, err
	}

	return leaveType, nil
}

// ListLeaveTypes gets all leave types in your company.
func (c Client) ListLeaveTypes() ([]LeaveType, error) {
	var leaveTypes []LeaveType

	resp, err := c.get(leaveTypeURL, nil)
	if err != nil {
		return leaveTypes, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&leaveTypes); err != nil {
		return leaveTypes, err
	}

	return leaveTypes, nil
}

// UpdateLeaveType update the given leave type id with the given
// request data
func (c Client) UpdateLeaveType(id string, lt UpdateLeaveTypeRequest) (LeaveType, error) {
	var leaveType LeaveType

	bytes, err := json.Marshal(lt)
	if err != nil {
		return leaveType, err
	}

	resp, err := c.put(leaveTypeURL+"/"+id, bytes)
	if err != nil {
		return leaveType, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&leaveType); err != nil {
		return leaveType, err
	}

	return leaveType, nil
}

// Leave contains all the leave information
type Leave struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	EmployeeID  int    `json:"employee_id"`
	FinishOn    string `json:"finish_on"`
	HalfDay     string `json:"half_day"`
	LeaveTypeID int    `json:"leave_type_id"`
	StartOn     string `json:"start_on"`
}

// CreateLeaveRequest keeps the information needed
// for create a new leave
type CreateLeaveRequest struct {
	Description string `json:"description,omitempty"`
	EmployeeID  int    `json:"employee_id"`
	FinishOn    string `json:"finish_on"`
	HalfDay     string `json:"half_day,omitempty"`
	LeaveTypeID int    `json:"leave_type_id"`
	StartOn     string `json:"start_on"`
}

// UpdateLeaveRequest keeps the information needed
// for update a leave
type UpdateLeaveRequest struct {
	Description string `json:"description,omitempty"`
	EmployeeID  int    `json:"employee_id,omitempty"`
	FinishOn    string `json:"finish_on,omitempty"`
	HalfDay     string `json:"half_day,omitempty"`
	LeaveTypeID int    `json:"leave_type_id,omitempty"`
	StartOn     string `json:"start_on,omitempty"`
}

// CreateLeave creates a new leave.
// Admins can create leaves for all employees,
// regular users are restricted to themselves and employees they manage.
func (c Client) CreateLeave(l CreateLeaveRequest) (Leave, error) {
	var leave Leave

	bytes, err := json.Marshal(l)
	if err != nil {
		return leave, err
	}

	resp, err := c.post(leaveURL, bytes)
	if err != nil {
		return leave, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&leave); err != nil {
		return leave, err
	}

	return leave, nil
}

// DeleteLeave deletes an existing leave.
// Admins can delete leaves for all employees,
// regular users are restricted to themselves and employees they manage.
// Restrictions apply if the leave has already started.
func (c Client) DeleteLeave(id string) error {
	_, err := c.delete(leaveURL + "/" + id)
	if err != nil {
		return err
	}

	return nil
}

// ListLeaves gets all leaves from your company.
func (c Client) ListLeaves() ([]Leave, error) {
	var leaves []Leave

	resp, err := c.get(leaveURL, nil)
	if err != nil {
		return leaves, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&leaves); err != nil {
		return leaves, err
	}

	return leaves, nil
}

// UpdateLeave update the given leave id with the given request data.
// Admins can update leaves for all employees,
// regular users are restricted to themselves and employees they manage.
// Restrictions apply if the leave has already started.
func (c Client) UpdateLeave(id string, lt UpdateLeaveRequest) (Leave, error) {
	var leave Leave

	bytes, err := json.Marshal(lt)
	if err != nil {
		return leave, err
	}

	resp, err := c.put(leaveURL+"/"+id, bytes)
	if err != nil {
		return leave, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&leave); err != nil {
		return leave, err
	}

	return leave, nil
}
