package kickertool

import "fmt"

const getDisciplineEntriesURL = "https://api.tournament.io/v1/public/tournaments/%s/discipline/%s/entries"

// Entry .
type Entry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetDisciplineEntries .
func (api API) GetDisciplineEntries(tournamentID, disciplineID string) ([]Entry, error) {
	url := fmt.Sprintf(getDisciplineEntriesURL, tournamentID, disciplineID)
	var response []Entry
	err := api.DoAPIRequest("GET", url, nil, nil, &response)
	if err != nil {
		return nil, nil
	}

	return response, nil
}
