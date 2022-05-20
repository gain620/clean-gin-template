package middleware

import (
	"clean-gin-template/pkg/logger"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Middleware -.
type Middleware interface {
	CORS() gin.HandlerFunc
}

type ginMiddleware struct {
	l logger.Interface
}

// New creates a middleware
func New(l logger.Interface) Middleware {
	return &ginMiddleware{
		l: l,
	}
}

// CORS will handle the CORS middleware
func (m *ginMiddleware) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PUT , UPDATE")

		log.WithFields(log.Fields{
			"middleware": "CORS",
		}).Info("MIDDLEWARE RUNNING...")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
