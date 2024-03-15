package model

import "time"

type Score struct {
	Code      uint      `gorm:"primaryKey;index"    json:"code"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	Name      string    `gorm:"not null"            json:"name"`
	Count     uint      `gorm:"not null"            json:"count"`
}

func (m Score) TableName() string {
	return "scores"
}
