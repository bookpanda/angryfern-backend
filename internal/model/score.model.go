package model

type Score struct {
	Base
	Code        string `json:"code" gorm:"primaryKey;index"`
	CountryName string `json:"country_name" gorm:"mediumtext"`
	ClickCount  uint   `json:"click_count" gorm:"type:bigint"`
}

func (m Score) TableName() string {
	return "scores"
}
