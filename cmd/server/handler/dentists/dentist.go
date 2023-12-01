package dentists

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ncondezo/final/internal/dentists"
	"github.com/ncondezo/final/internal/domain"
	"github.com/ncondezo/final/pkg/web"
)

type Controller struct {
	service dentists.Service
}

func NewDentistController(service dentists.Service) *Controller {
	return &Controller{service: service}
}

// @BasePath /api/v1

// HandlerCreate godoc
// @Summary Create a new dentist
// @Tags dentists
// @Accept json
// @Produce json
// @Param Dentist body domain.DentistDTO true "Dentist information"
// @Success 201 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /dentists [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.DentistDTO

		err := ctx.Bind(&request)

		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "bad request")
			return
		}

		dentist, err := c.service.Create(ctx, request)
		if errors.Is(err, dentists.ErrAlreadyExists) {
			web.NewErrorResponse(ctx, http.StatusConflict, "dentist already exists")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusCreated, dentist)
	}
}

// HandlerGetByID godoc
// @Summary Get a dentist by id
// @Tags dentists
// @Produce json
// @Param ID path int true "Dentist ID to search"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /dentists/:id [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "invalid id")
			return
		}

		dentist, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, dentist)
	}
}

// HandlerUpdate godoc
// @Summary Update a dentist by id
// @Tags dentists
// @Accept json
// @Produce json
// @Param ID path int true "Dentist ID to update"
// @Param Dentist body domain.DentistDTO true "Dentist information"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /dentists/:id [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.DentistDTO

		errBind := ctx.Bind(&request)
		if errBind != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "bad request binding")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "bad request param")
			return
		}

		dentist, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, dentist)
	}
}

// HandlerPatch godoc
// @Summary Update a dentist registry number by id
// @Tags dentists
// @Accept json
// @Produce json
// @Param ID path int true "Dentist ID to update"
// @Param DentistRegistry body domain.DentistRegistryDTO true "Dentist registry number"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /dentists/:id [patch]
func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "invalid id")
			return
		}

		var request domain.DentistRegistryDTO

		errBind := ctx.Bind(&request)
		if errBind != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "bad request binding")
			return
		}

		dentist, err := c.service.Patch(ctx, request, id)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, dentist)
	}
}

// HandlerDelete godoc
// @Summary Delete a dentist by id
// @Tags dentists
// @Accept json
// @Produce json
// @Param ID path int true "Dentist ID to delete"
// @Success 200 {object} web.SuccessResponse
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /dentists/:id [delete]
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "invalid id")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, gin.H{
			"message": "dentist deleted",
		})
	}
}
