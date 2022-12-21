package models

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	Poll    Poll   `json:"-"`
	Target  Target `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   bool   `json:"err"`
}
