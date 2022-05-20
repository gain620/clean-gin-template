package v1

import (
	"clean-gin-template/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	//_ "github.com/evrone/go-clean-template/docs"
	"clean-gin-template/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Gin Clean Template API
// @description Using a github web api service as an example
// @version     1.0
// @host        localhost:{PORT}
// @BasePath    /v1
func NewRouter(handler *gin.Engine, g usecase.Github, l logger.Interface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"timestamp": time.Now().Unix(),
		})
	})

	// Prometheus metrics
	//handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		NewGithubRoutes(h, g, l)
	}
}
