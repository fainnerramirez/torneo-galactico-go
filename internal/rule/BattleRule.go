package rule

import (
	"strings"

	"api-torneo/internal/domain"
)

func DetermineWinner(specieA domain.Specie, specieB domain.Specie) domain.Specie {
	if specieA.PowerLevel != specieB.PowerLevel {
		if specieA.PowerLevel > specieB.PowerLevel {
			return specieA
		}
		return specieB
	}

	if strings.Compare(strings.ToLower(specieA.Name), strings.ToLower(specieB.Name)) <= 0 {
		return specieA
	}
	return specieB
}
