package middleware

import (
	"clean-gin-template/pkg/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Middleware -.
type Middleware interface {
	JWT() gin.HandlerFunc
	Limiter(*limiter.Limiter) gin.HandlerFunc
	CORS() gin.HandlerFunc
}

type ginMiddleware struct {
	l logger.Interface
}

// JWT middleware validates the token from the http request, returning a 401 if it's not valid
func (m *ginMiddleware) JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		//token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			m.l.Print("Claims[Name]: ", claims["name"])
			m.l.Print("Claims[Admin]: ", claims["admin"])
			m.l.Print("Claims[Issuer]: ", claims["iss"])
			m.l.Print("Claims[IssuedAt]: ", claims["iat"])
			m.l.Print("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			m.l.Print(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
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
