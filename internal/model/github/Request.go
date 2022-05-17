package model

type ContributorRequest struct {
	Owner string `uri:"owner" form:"owner" json:"owner" binding:"required"`
	Repo  string `uri:"repo" form:"repo" json:"repo" binding:"required"`
}

// ErrorMessage defines error body struct
type ErrorMessage struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
