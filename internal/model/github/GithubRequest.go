package model

type Request struct {
	Owner string
	Repo  string
}

// ErrorMessage defines error body struct
type ErrorMessage struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
