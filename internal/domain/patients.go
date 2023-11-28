package domain

import "time"

type Patient struct {
	Name             string    `json:"name"`
	LastName         string    `json:"last_name"`
	Address          string    `json:"address"`
	Document         string    `json:"document"`
	RegistrationDate time.Time `json:"registration_date"`
}
