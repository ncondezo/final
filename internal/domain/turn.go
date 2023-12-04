package domain

import "time"

type Turn struct {
	Id          int       `json:"id"`
	IdDentist   int       `json:"id_dentist"`
	IdPatient   int       `json:"id_patient"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

type TurnDTO struct {
	IdDentist   int       `json:"id_dentist"`
	IdPatient   int       `json:"id_patient"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}
