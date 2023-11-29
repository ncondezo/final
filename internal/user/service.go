package product

import (
	"errors"
	"log"

	"github.com/ncondezo/final/internal/domain"
	"github.com/ncondezo/final/pkg/security"
	"github.com/ncondezo/final/pkg/web"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorInvalidCredentials = errors.New("invalid credentials")
)

type Service interface {
	Signup(dto domain.SignupDTO) (*domain.User, error)
	Login(dto domain.LoginDTO) (*web.LoginResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (service *service) Signup(dto domain.SignupDTO) (*domain.User, error) {
	passwordEncrypted, err := passwordEncrypt(dto.Password)
	if err != nil {
		return nil, err
	}
	userData := domain.User{
		uuid.New().String(),
		dto.Name,
		dto.Surname,
		dto.Email,
		passwordEncrypted,
	}
	return service.repository.Create(&userData)
}

func (service *service) Login(dto domain.LoginDTO) (*web.LoginResponse, error) {
	user, err := service.repository.FindByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	if !passwordCompare(dto.Password, user.Password) {
		return nil, ErrorInvalidCredentials
	}
	token, err := security.GenerateToken(&dto)
	if err != nil {
		return nil, err
	}
	return &web.LoginResponse{token}, nil
}

func passwordEncrypt(password string) (string, error) {
	passwordBytes := []byte(password)
	log.Println(passwordBytes)
	encrypted, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

func passwordCompare(password, hashed string) bool {
	passwordBytes := []byte(password)
	hashedBytes := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(hashedBytes, passwordBytes)
	return err == nil
}
