package kickertool

import "fmt"

const getTournamentURL = "https://api.tournament.io/v1/public/tournaments/%s"

// Discipline .
type Discipline struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	EntryType string `json:"entryType"`
}

// Tournament .
type Tournament struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Disciplines []Discipline `json:"disciplines"`
}

// GetTournamentDisciplines .
func (api API) GetTournament(tournamentID string) (*Tournament, error) {
	url := fmt.Sprintf(getTournamentURL, tournamentID)
	var response Tournament
	err := api.DoAPIRequest("GET", url, nil, nil, &response)
	if err != nil {
		return nil, nil
	}

	return &response, nil
}
