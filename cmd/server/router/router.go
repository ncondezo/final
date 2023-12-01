package router

import (
	"database/sql"
	"net/http"

	authController "github.com/ncondezo/final/cmd/server/handler/auth"
	patientController "github.com/ncondezo/final/cmd/server/handler/patient"
	patient "github.com/ncondezo/final/internal/patients"
	user "github.com/ncondezo/final/internal/user"
	"github.com/ncondezo/final/pkg/middleware"
	_ "github.com/ncondezo/final/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Routes interface {
	BuildRoutes()
}

type router struct {
	engine   *gin.Engine
	apiGroup *gin.RouterGroup
	db       *sql.DB
}

func NewRouter(engine *gin.Engine, db *sql.DB) Routes {
	return &router{engine: engine, db: db}
}

func (router *router) BuildRoutes() {
	router.setApiGroup()
	router.buildPingEndpoint()
	router.buildSwaggerEndpoint()
	router.buildAuthGroup()
	router.buildPatients()
}

func (router *router) setApiGroup() {
	router.apiGroup = router.engine.Group("/api/v1")
}

func (router *router) buildPingEndpoint() {
	router.apiGroup.GET("/health",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"status": "up"})
		})
}

func (router *router) buildSwaggerEndpoint() {
	router.apiGroup.GET("/docs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (router *router) buildAuthGroup() {

	repository := user.NewRepository(router.db)
	service := user.NewService(repository)
	controller := authController.NewController(service)

	authGroup := router.apiGroup.Group("/auth")
	authGroup.POST("/signup", controller.Signup())
	authGroup.POST("/login", controller.Login())

}

func (router *router) buildPatients() {

	repository := patient.NewMyRepository(router.db)
	service := patient.NewServicePatient(repository)
	controller := patientController.NewControladorPaciente(service)

	patientGroup := router.apiGroup.Group("/patients")
	{
		patientGroup.POST("", middleware.Authorization(), controller.HandlerCreate())
		patientGroup.GET("/:id", middleware.Authorization(), controller.HandlerGetByID())
		patientGroup.PUT("/:id", middleware.Authorization(), controller.HandlerUpdate())
		patientGroup.PATCH("/:id", middleware.Authorization(), controller.HandlerPatch())
		patientGroup.DELETE("/:id", middleware.Authorization(), controller.HandlerDelete())
	}

}
