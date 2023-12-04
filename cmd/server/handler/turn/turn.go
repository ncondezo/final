package turn

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ncondezo/final/internal/domain"
	"github.com/ncondezo/final/internal/turns"
	"github.com/ncondezo/final/pkg/web"
)

type Controller struct {
	service turns.Service
}

func NewTurnController(service turns.Service) *Controller {
	return &Controller{service: service}
}

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
		if errors.Is(err, turns.ErrAlreadyExists) {
			web.NewErrorResponse(ctx, http.StatusConflict, "turn already exists")
			return
		}
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusCreated, turn)
	}
}

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
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.NewSuccessResponse(ctx, http.StatusOK, turn)
	}
}

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
