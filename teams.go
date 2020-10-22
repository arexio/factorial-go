package factorial

import "encoding/json"

const (
	teamURL = "/api/v1/teams"
)

// Team type holds the basic information for a team
// defined from Factorial
type Team struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	EmployeeIDs []int  `json:"employee_ids"`
	LeadIDs     []int  `json:"lead_ids"`
}

// GetTeam will get the team by the given id
func (c Client) GetTeam(id string) (Team, error) {
	var team Team

	resp, err := c.get(teamURL+"/"+id, nil)
	if err != nil {
		return team, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&team); err != nil {
		return team, err
	}

	return team, nil
}

// ListTeams will get all the teams saved in Factorial
func (c Client) ListTeams() ([]Team, error) {
	var teams []Team

	resp, err := c.get(teamURL, nil)
	if err != nil {
		return teams, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&teams); err != nil {
		return teams, err
	}

	return teams, nil
}
