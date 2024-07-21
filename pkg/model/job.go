package model

import (
	"time"

	"gorm.io/gorm"
)



type SqlJob struct {
	gorm.Model
	ID               uint      `json:"id" gorm:"not null;primaryKey;auto_increment"`
	Name             string    `json:"name" gorm:"not null;unique;size:255"`
	Sql              string    `json:"sql" gorm:"not null"`
	Status           string    `json:"status" gorm:"not null;size:255"`
	Frequency        string       `json:"frequency" gorm:"not null"`
	StartTime        time.Time `json:"starttime" gorm:"not null"`
	EndTime          time.Time `json:"endtime" gorm:"not null"`
	WechatGroup      string    `json:"wechatgroup" gorm:"not null;size:255"`
	Level            string    `json:"level" gorm:"not null;size:255"`
	DesensitiseOrNot bool      `json:"desensitiseornot" gorm:"not null;default:true"`
}

type ScriptJob struct {
	gorm.Model
	ID               uint      `json:"id" gorm:"not null;primaryKey;auto_increment"`
	Name             string    `json:"name" gorm:"not null;unique;size:255"`
	Script           string    `json:"script" gorm:"not null"`
	Status           string    `json:"status" gorm:"not null;size:255"`
	Frequency        string       `json:"frequency" gorm:"not null"`
	StartTime        time.Time `json:"starttime" gorm:"not null"`
	EndTime          time.Time `json:"endtime" gorm:"not null"`
	WechatGroup      string    `json:"wechatgroup" gorm:"not null;size:255"`
	Level            string    `json:"level" gorm:"not null;size:255"`
	DesensitiseOrNot bool      `json:"desensitiseornot" gorm:"not null;default:true"`
}
