package models

type ResponseSuccess struct {
	Data any `json:"data"`
}

type ResponseError struct {
	Error any `json:"error"`
}

type SortOrder int

const (
	Ascending  SortOrder = 1
	Descending SortOrder = -1
)
