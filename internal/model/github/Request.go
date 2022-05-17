package model

type ContributorRequest struct {
	Owner string `form:"owner" json:"owner" binding:"required"`
	Repo  string `form:"repo" json:"repo" binding:"required"`
}

// ErrorMessage defines error body struct
type ErrorMessage struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
