package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
)

// TODO: Implement Custom Error Wrapping
type response struct {
	ErrorCode int   `json:"code" example:"502"`
	ErrorMsg  error `json:"error" example:"message"`
}

//func errorResponse(c *gin.Context, customError CustomError) {
//	c.AbortWithStatusJSON(customError.Code, response{customError.Message})
//}

func errorResponse(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, response{ErrorCode: code, ErrorMsg: err})
}
