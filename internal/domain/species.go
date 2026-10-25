package domain

type Specie struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PowerLevel     int    `json:"powerlevel"`
	SpecialAbility string `json:"specialability"`
}
