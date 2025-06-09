package journeyModels

import (
	"time"

	"gorm.io/gorm"
)

type ExperienceType struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Experiences []Experience
}

type Experience struct {
	gorm.Model
	Title            string     `gorm:"not null"`
	Company          string     `gorm:"not null"`
	From             time.Time  `gorm:"not null"`
	To               *time.Time `json:"to,omitempty"`
	Diploma          *string    `json:"diploma,omitempty"`
	Tasks            []Task     `json:"tasks,omitempty"`
	ExperienceTypeId uint
	ExperienceType   ExperienceType
}

type Task struct {
	gorm.Model
	Details      string `gorm:"not null"`
	ExperienceId uint
}
