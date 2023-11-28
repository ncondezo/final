package domain

import "github.com/golang-jwt/jwt"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupDTO struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
