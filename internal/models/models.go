package models

type Link string
type Path string

type Model struct {
	Field  string `json:"field"`
	Field2 int    `json:"field2,omitempty"`
}
