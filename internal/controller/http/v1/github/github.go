package v1

import (
	model "clean-gin-template/internal/model/github"
	"net/http"

	"github.com/gin-gonic/gin"

	"clean-gin-template/internal/usecase"
	"clean-gin-template/pkg/logger"
)

type githubRoutes struct {
	g usecase.Github
	l logger.Interface
}

func NewGithubRoutes(handler *gin.RouterGroup, g usecase.Github, l logger.Interface) {
	r := &githubRoutes{g, l}

	h := handler.Group("/github")
	{
		h.GET("/contributors", r)
		//h.POST("/do-translate", r.doTranslate)
	}
}

type contributorsResponse struct {
	Contributors []model.Contributor `json:"contributors"`
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
func (r *githubRoutes) getContributors(c *gin.Context) {
	contributors, err := r.g.GetContributors(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "github web api problems")
		return
	}

	c.JSON(http.StatusOK, contributorsResponse{contributors})
}

//type doTranslateRequest struct {
//	Source      string `json:"source"       binding:"required"  example:"auto"`
//	Destination string `json:"destination"  binding:"required"  example:"en"`
//	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
//}
//
//// @Summary     Translate
//// @Description Translate a text
//// @ID          do-translate
//// @Tags  	    translation
//// @Accept      json
//// @Produce     json
//// @Param       request body doTranslateRequest true "Set up translation"
//// @Success     200 {object} entity.Translation
//// @Failure     400 {object} response
//// @Failure     500 {object} response
//// @Router      /translation/do-translate [post]
//func (r *translationRoutes) doTranslate(c *gin.Context) {
//	var request doTranslateRequest
//	if err := c.ShouldBindJSON(&request); err != nil {
//		r.l.Error(err, "http - v1 - doTranslate")
//		errorResponse(c, http.StatusBadRequest, "invalid request body")
//
//		return
//	}
//
//	translation, err := r.t.Translate(
//		c.Request.Context(),
//		entity.Translation{
//			Source:      request.Source,
//			Destination: request.Destination,
//			Original:    request.Original,
//		},
//	)
//	if err != nil {
//		r.l.Error(err, "http - v1 - doTranslate")
//		errorResponse(c, http.StatusInternalServerError, "translation service problems")
//
//		return
//	}
//
//	c.JSON(http.StatusOK, translation)
//}
