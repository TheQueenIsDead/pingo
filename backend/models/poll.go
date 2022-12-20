package models

import (
	"gorm.io/gorm"
	"time"
)

type Poll struct {
	gorm.Model
	Targets         []int     `json:"targets"`
	DurationSeconds int       `json:"duration"`
	StartedAt       time.Time `json:"started"`
}
