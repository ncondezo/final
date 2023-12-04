package domain

import "time"

type Turn struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Patient     Patient   `json:"patient"`
	Dentist     Dentist   `json:"dentist"`
}

type TurnDTO struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	IdPatient   int       `json:"id_patient"`
	IdDentist   int       `json:"id_dentist"`
}
