package api

import "gorm.io/gorm"

type Application struct {
	db *gorm.DB
}

func NewApplication(db *gorm.DB) *Application {
	return &Application{db: db}
}
