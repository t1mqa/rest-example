package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rest-example/internal/cfg"
	"rest-example/internal/models"
)

type GormDatabase struct {
	db *gorm.DB
}

func NewGormDatabase(db *gorm.DB) Database {
	return &GormDatabase{
		db: db,
	}
}

func (gdb *GormDatabase) GetByID(id uint) (*models.Model, error) {
	var model models.Model
	result := gdb.db.First(&model, id)
	return &model, result.Error
}

func NewDB(appConfig cfg.DBSettings) Database {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appConfig.Host, appConfig.Port, appConfig.User, appConfig.Password, appConfig.DBName)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Model{})
	if err != nil {
		panic(err)
	}

	return NewGormDatabase(db)
}
