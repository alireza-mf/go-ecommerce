package models

type ResponseSuccess struct {
	Data any `json:"data"`
}

type ResponseError struct {
	Error any `json:"error"`
}
