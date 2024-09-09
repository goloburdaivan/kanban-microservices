package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255)"`
	Description string `json:"description"`
}
