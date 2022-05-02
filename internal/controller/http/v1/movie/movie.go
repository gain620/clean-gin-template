// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	//_ "github.com/evrone/go-clean-template/docs"
	"clean-gin-template/pkg/logger"
)

// NewMovieRouter -.
// Swagger spec:
// @title       Gin Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewMovieRouter(handler *gin.Engine, l logger.Interface) {
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
	//h := handler.Group("/v1")
	//{
	//	//newTranslationRoutes(h, t, l)
	//}
}
