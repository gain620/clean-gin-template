package middleware

import (
	"clean-gin-template/pkg/logger"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
)

// Middleware -.
type Middleware interface {
	Limiter(*limiter.Limiter) gin.HandlerFunc
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

// Limiter will handle the rate limiting
func (m *ginMiddleware) Limiter(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// CORS will handle the CORS middleware
func (m *ginMiddleware) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PUT , UPDATE")

		m.l.Debug("Middleware running")

		//if c.Request.Method == "OPTIONS" {
		//	c.AbortWithStatus(204)
		//	return
		//}

		c.Next()
	}
}
