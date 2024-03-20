package storage

import (
	"rest-example/internal/models"
)

type Database interface {
	GetByID(id uint) (*models.Model, error)
}
