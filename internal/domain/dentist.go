package domain

type Dentist struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"lastname"`
	Registration string `json:"registry"`
}

type DentistDTO struct {
	Name         string `json:"name"`
	LastName     string `json:"lastname"`
	Registration string `json:"registry"`
}

type DentistRegistryDTO struct {
	Registration string `json:"registry"`
}
