package models

import (
	"gorm.io/gorm"
	"time"
)

type Poll struct {
	gorm.Model
	TargetIds       []int     `json:"target_ids,omitempty" gorm:"-"`                               // Skip persisting to db, but use in JSON for transfer
	Targets         []Target  `json:"-" gorm:"many2many:poll_targets;foreignKey:ID;references:ID"` // Skip representing Target objects in JSON, but use foreign keys
	DurationSeconds int       `json:"duration"`
	StartedAt       time.Time `json:"started"`
}
