package models

import (
	"gorm.io/gorm"
	"time"
)

// {"time":"2022-12-20T22:13:07.7005205+13:00","level":"ERROR","prefix":"-","file":"poll.go","line":"18","message":"
// code=400, message=
// Unmarshal type error: expected=models.Target,
// got=number, field=target_ids, offset=23,
// internal=json: cannot unmarshal number into Go struct field Poll.target_ids of type models.
type Poll struct {
	gorm.Model
	TargetIds       []Target  `json:"target_ids" gorm:"-"`
	Targets         []Target  `json:"" gorm:"foreignKey:ID"`
	DurationSeconds int       `json:"duration"`
	StartedAt       time.Time `json:"started"`
}
