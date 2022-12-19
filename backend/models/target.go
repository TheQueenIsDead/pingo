package models

import "gorm.io/gorm"

type Target struct {
	gorm.Model
	Type      string `json:"type"`
	Source    string `json:"source"`
	Frequency int    `json:"frequency"`
	Unit      string `json:"unit"`
}
