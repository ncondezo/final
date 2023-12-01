package user

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/ncondezo/final/internal/domain"
	user "github.com/ncondezo/final/internal/user"
	"github.com/ncondezo/final/pkg/web"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service user.Service
}

func NewController(service user.Service) *controller {
	return &controller{service}
}

// @BasePath /api/v1

// Signup godoc
// @Summary Register a new user
// @Description Takes user information and store in DB. Return saved user.
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.SignupDTO true "User register information"
// @Success 201 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /auth/signup [post]
func (controller *controller) Signup() gin.HandlerFunc {
	return func(context *gin.Context) {
		var userData domain.SignupDTO
		err := context.ShouldBindJSON(&userData)
		if err != nil {
			web.NewErrorResponse(context, http.StatusBadRequest,
				"El JSON enviado en el cuerpo no es válido")
			return
		}
		if err := reqJsonValidation(userData); err != "" {
			web.NewErrorResponse(context, http.StatusBadRequest, err)
			return
		}
		created, err := controller.service.Signup(userData)
		if errors.Is(err, user.ErrorUserExists) {
			web.NewErrorResponse(context, http.StatusConflict,
				"El usuario con email "+userData.Email+" ya existe")
			return
		}
		if err != nil {
			web.NewErrorResponse(context, http.StatusInternalServerError,
				"Se ha producido un error al crear el usuario")
			return
		}
		web.NewSuccessResponse(context, http.StatusCreated, created)
	}
}

// Login godoc
// @Summary Existing user login
// @Description Takes and verify user credentials. Returns an access token for the user.
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.LoginDTO true "User credentials"
// @Success 200 {object} web.LoginResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /auth/login [post]
func (controller *controller) Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		var userData domain.LoginDTO
		err := context.ShouldBindJSON(&userData)
		if err != nil {
			web.NewErrorResponse(context, http.StatusBadRequest,
				"El JSON enviado en el cuerpo no es válido")
			return
		}
		if err := reqJsonValidation(userData); err != "" {
			web.NewErrorResponse(context, http.StatusBadRequest, err)
			return
		}
		logged, err := controller.service.Login(userData)
		if errors.Is(err, user.ErrorUserNotFound) {
			web.NewErrorResponse(context, http.StatusNotFound,
				"El usuario con email "+userData.Email+" no existe")
			return
		}
		if errors.Is(err, user.ErrorInvalidCredentials) {
			web.NewErrorResponse(context, http.StatusForbidden,
				"Las credenciales son inválidas")
			return
		}
		if err != nil {
			web.NewErrorResponse(context, http.StatusInternalServerError,
				"Se ha producido un error al intentar loguear el usuario")
			return
		}
		web.NewLoginResponse(context, http.StatusOK, *logged)
	}
}

func reqJsonValidation(request interface{}) string {
	requestCamps := reflect.ValueOf(request)
	for i := 0; i < requestCamps.NumField(); i++ {
		campName := requestCamps.Type().Field(i).Name
		campValue := requestCamps.Field(i).Interface()
		campType := fmt.Sprint(reflect.TypeOf(campValue).Kind())
		switch campType {
		case "string":
			if campValue == "" {
				return fmt.Sprintf("El campo %v no puede estar vacío", campName)
			}
		case "int":
			if campValue.(int) <= 0 {
				return fmt.Sprintf("El campo %v no puede estar vacío", campName)
			}
		case "float64":
			if campValue.(float64) == 0 {
				return fmt.Sprintf("El campo %v no puede estar vacío", campName)
			}
		case "boolean":
			if !campValue.(bool) {
				return fmt.Sprintf("El campo %v no puede estar vacío", campName)
			}

		}
	}
	return ""
}
