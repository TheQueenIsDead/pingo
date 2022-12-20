package models

import "gorm.io/gorm"

type Target struct {
	gorm.Model
	Source    string `json:"source"`
	Frequency int    `json:"frequency"`
	Unit      string `json:"unit"`
}
