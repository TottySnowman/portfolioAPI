package tagModel

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Tag  string `gorm:"size:20"`          // Name of the tag
	Icon string `gorm:"size:20;not null"` // Icon representing the tag
}

type JsonTag struct {
	TagId int
	Icon  string
	Tag   string
}
