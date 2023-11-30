package patient

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ncondezo/final/internal/domain"
	"github.com/ncondezo/final/internal/patients"
	"github.com/ncondezo/final/pkg/web"
)

type Controlador struct {
	service patients.Service
}

func NewControladorPaciente(service patients.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

func (c *Controlador) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Patient

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		patient, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, patient)

	}
}

// Producto godoc
// @Summary producto example
// @Description Get producto by id
// @Tags producto
// @Param id path int true "id del producto"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /productos/:id [get]
func (c *Controlador) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		patient, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, patient)
	}
}

// Producto godoc
// @Summary producto example
// @Description Update producto by id
// @Tags producto
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /productos/:id [put]
func (c *Controlador) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Patient

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		patient, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, patient)

	}
}

// Producto godoc
// @Summary producto example
// @Description Delete producto by id
// @Tags producto
// @Param id path int true "id del producto"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /productos/:id [delete]
func (c *Controlador) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, "producto eliminado")
	}
}

// Producto godoc
// @Summary producto example
// @Description Patch producto
// @Tags producto
// @Param id path int true "id del producto"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /productos/:id [patch]
func (c *Controlador) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		var request domain.Patient

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		patient, err := c.service.Patch(ctx, request, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, patient)
	}
}
