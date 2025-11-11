package tagModel

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Tag  string `gorm:"size:20"`          // Name of the tag
	Icon string `gorm:"size:20;not null"` // Icon representing the tag
	IsBackend bool `gorm:"not null"`
}

type JsonTag struct {
	TagId int
	Icon  string
	Tag   string
}

type TechstackItem struct{
	TagId int
	Icon  string
	Tag   string
  IsBackend bool
}

type TechstackItemResponse struct{
	IconName  string `json:"iconName"`
	DisplayName   string `json:"displayName"`
}

type TechstackResponse struct{
  Frontend []TechstackItemResponse `json:"frontend"`
  Backend []TechstackItemResponse`json:"backend"`
}
