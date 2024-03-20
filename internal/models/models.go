package models

type Link string
type Path string

type Model struct {
	Field  string `json:"field" gorm:"type:varchar(32);not null'"`
	Field2 int    `json:"field2,omitempty" gorm:"default:111"`
}
