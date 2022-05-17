package v1

import "github.com/gin-gonic/gin"

// TODO: Implement Custom Error Wrapping
type response struct {
	ErrorCode int    `json:"code" example:"502"`
	ErrorMsg  string `json:"error" example:"message"`
}

//func errorResponse(c *gin.Context, customError CustomError) {
//	c.AbortWithStatusJSON(customError.Code, response{customError.Message})
//}

func errorResponse(c *gin.Context, code int, err string) {
	c.AbortWithStatusJSON(code, response{ErrorCode: code, ErrorMsg: err})
}
