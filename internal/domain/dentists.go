package domain

type Dentist struct {
	Id           int    `json:"id"`
	LastName     string `json:"last_name"`
	Name         string `json:"name"`
	Registration string `json:"registration"`
}
