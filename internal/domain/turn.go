package domain

import "time"

type Turn struct {
	Id        int       `json:"id"`
	IdDentist int       `json:"id_dentist"`
	IdPatient int       `json:"id_patient"`
	DateUp    time.Time `json:"dateup"`
}
