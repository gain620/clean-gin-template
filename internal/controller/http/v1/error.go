package v1

import "github.com/gin-gonic/gin"

// TODO: Implement Custom Error Wrapping
type response struct {
	Error string `json:"error" example:"message"`
}

type serverError struct {
	Code    int    `json:"code" example:"501"`
	Message string `json:"error" example:"error_message"`
}

//func errorResponse(c *gin.Context, customError CustomError) {
//	c.AbortWithStatusJSON(customError.Code, response{customError.Message})
//}

func errorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, response{message})
}
