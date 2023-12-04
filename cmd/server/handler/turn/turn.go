package turn

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ncondezo/final/internal/dentists"
	"github.com/ncondezo/final/internal/domain"
	"github.com/ncondezo/final/internal/patients"
	"github.com/ncondezo/final/internal/turns"
	"github.com/ncondezo/final/pkg/web"
)

type Controller struct {
	service turns.Service
}

func NewTurnController(service turns.Service) *Controller {
	return &Controller{service: service}
}

// @BasePath /api/v1

// HandlerCreate godoc
// @Summary Create a new turn
// @Tags turns
// @Accept json
// @Produce json
// @Param Turn body domain.TurnDTO true "Turn information"
// @Success 201 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /turns [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.TurnDTO

		err := ctx.Bind(&request)

		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "bad request")
			return
		}
		if err := web.RequestJsonValidation(request); err != "" {
			web.NewErrorResponse(ctx, http.StatusBadRequest, err)
			return
		}

		turn, err := c.service.Create(ctx, request)
		if errors.Is(err, patients.ErrNotFound) {
			web.NewErrorResponse(ctx, http.StatusNotFound, "patient not found")
			return
		}
		if errors.Is(err, dentists.ErrNotFound) {
			web.NewErrorResponse(ctx, http.StatusNotFound, "dentist not found")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusCreated, turn)
	}
}

// HandlerGetByID godoc
// @Summary Get a turn by id
// @Tags turns
// @Produce json
// @Param ID path int true "Turn ID to search"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /turns/:id [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "invalid id")
			return
		}

		turn, err := c.service.GetByID(ctx, id)
		if errors.Is(err, turns.ErrNotFound) {
			web.NewErrorResponse(ctx, http.StatusNotFound, "turn not found")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, turn)
	}
}

// HandlerGetByPatientID godoc
// @Summary Get a turn by patient id
// @Tags turns
// @Produce json
// @Param ID path int true "Patient ID to search"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /turns/patient/:id [get]
func (c *Controller) HandlerGetByPatientID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		patientId, err := strconv.Atoi(ctx.Param("patientId"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "invalid patient id")
			return
		}

		turn, err := c.service.GetByPatientID(ctx, patientId)
		if errors.Is(err, patients.ErrNotFound) {
			web.NewErrorResponse(ctx, http.StatusNotFound, "patient not found")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, turn)
	}
}

// HandlerUpdate godoc
// @Summary Update a turn by id
// @Tags turns
// @Accept json
// @Produce json
// @Param ID path int true "Turn ID to update"
// @Param Turn body domain.TurnDTO true "Turn information"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /turns/:id [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.TurnDTO

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

		turn, err := c.service.Update(ctx, request, id)
		if errors.Is(err, turns.ErrNotFound) {
			web.NewErrorResponse(ctx, http.StatusNotFound, "turn not found")
			return
		}
		if errors.Is(err, dentists.ErrNotFound) {
			web.NewErrorResponse(ctx, http.StatusNotFound, "dentist not found")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, turn)
	}
}

// HandlerDelete godoc
// @Summary Delete a turn by id
// @Tags turns
// @Accept json
// @Produce json
// @Param ID path int true "Turn ID to delete"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /turns/:id [delete]
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "invalid id")
			return
		}

		err = c.service.Delete(ctx, id)
		if errors.Is(err, turns.ErrNotFound) {
			web.NewErrorResponse(ctx, http.StatusNotFound, "turn not found")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, gin.H{
			"message": "turn deleted",
		})
	}
}
