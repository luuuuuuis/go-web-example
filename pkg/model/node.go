package model

import (
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null;unique;size:255"`
	IP     string `json:"ip" gorm:"not null;unique;size:255"`
	Port   int    `json:"port" gorm:"not null"`
	Status string `json:"status" gorm:"not null;size:255"`
	ID     int    `json:"id" gorm:"not null;primaryKey;auto_increment"`
}
