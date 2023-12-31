package patient

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ncondezo/final/internal/domain"
	"github.com/ncondezo/final/internal/patients"
	"github.com/ncondezo/final/pkg/web"
)

type Controller struct {
	service patients.Service
}

func NewPatientController(service patients.Service) *Controller {
	return &Controller{service: service}
}

// @BasePath /api/v1

// HandlerCreate godoc
// @Summary Create a new patient
// @Tags patients
// @Accept json
// @Produce json
// @Param Patient body domain.PatientDTO true "Patient information"
// @Success 201 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 409 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /patients [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.PatientDTO

		err := ctx.Bind(&request)

		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "bad request")
			return
		}
		if err := web.RequestJsonValidation(request); err != "" {
			web.NewErrorResponse(ctx, http.StatusBadRequest, err)
			return
		}

		patient, err := c.service.Create(ctx, request)
		if errors.Is(err, patients.ErrAlreadyExists) {
			web.NewErrorResponse(ctx, http.StatusConflict, "patient already exists")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusCreated, patient)
	}
}

// HandlerGetByID godoc
// @Summary Get a patient by id
// @Tags patients
// @Produce json
// @Param ID path int true "Patient ID to search"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 404 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /patients/:id [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "invalid id")
			return
		}

		patient, err := c.service.GetByID(ctx, id)
		if errors.Is(err, patients.ErrNotFound) {
			web.NewErrorResponse(ctx, http.StatusNotFound, "patient not found")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, patient)
	}
}

// HandlerUpdate godoc
// @Summary Update a patient by id
// @Tags patients
// @Accept json
// @Produce json
// @Param ID path int true "Patient ID to update"
// @Param Patient body domain.PatientDTO true "Patient information"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 404 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /patients/:id [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.PatientDTO

		errBind := ctx.Bind(&request)
		if errBind != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "bad request binding")
			return
		}
		if err := web.RequestJsonValidation(request); err != "" {
			web.NewErrorResponse(ctx, http.StatusBadRequest, err)
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "invalid id")
			return
		}

		patient, err := c.service.Update(ctx, request, id)
		if errors.Is(err, patients.ErrNotFound) {
			web.NewErrorResponse(ctx, http.StatusNotFound, "patient not found")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, patient)
	}
}

// HandlerPatch godoc
// @Summary Update a patient dni by id
// @Tags patients
// @Accept json
// @Produce json
// @Param ID path int true "Patient ID to update"
// @Param Patient body domain.PatientDniDTO true "Patient DNI"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 404 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /patients/:id [patch]
func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.PatientDniDTO

		errBind := ctx.Bind(&request)
		if errBind != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "bad request binding")
			return
		}
		if err := web.RequestJsonValidation(request); err != "" {
			web.NewErrorResponse(ctx, http.StatusBadRequest, err)
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "invalid id")
			return
		}

		patient, err := c.service.Patch(ctx, request, id)
		if errors.Is(err, patients.ErrNotFound) {
			web.NewErrorResponse(ctx, http.StatusNotFound, "patient not found")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, patient)
	}
}

// HandlerDelete godoc
// @Summary Delete a patient by id
// @Tags patients
// @Accept json
// @Produce json
// @Param ID path int true "Patient ID to delete"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 404 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /patients/:id [delete]
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "invalid id")
			return
		}

		err = c.service.Delete(ctx, id)
		if errors.Is(err, patients.ErrNotFound) {
			web.NewErrorResponse(ctx, http.StatusNotFound, "patient not found")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, gin.H{
			"message": "patient deleted",
		})
	}
}
