package domain

type Dentist struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"surname"`
	Registration string `json:"registry"`
}
