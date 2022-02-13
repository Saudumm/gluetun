package validation

import (
	"github.com/qdm12/gluetun/internal/models"
)

func ExpressvpnCountriesChoices(servers []models.ExpressvpnServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Country
	}
	return makeUnique(choices)
}

func ExpressvpnCityChoices(servers []models.ExpressvpnServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].City
	}
	return makeUnique(choices)
}

func ExpressvpnHostnameChoices(servers []models.ExpressvpnServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Hostname
	}
	return makeUnique(choices)
}
