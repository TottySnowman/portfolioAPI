package journeyModels

import (
	"time"

	"gorm.io/gorm"
)

type ExperienceType struct {
	gorm.Model
	Name string
}

type Experience struct {
	gorm.Model
	Title   string
	Company string
	From    time.Time
	To      *time.Time `json:"to,omitempty"`
	Diploma *string    `json:"diploma,omitempty"`
	Tasks   []Task     `json:"tasks,omitempty"`
	Type    ExperienceType
}

type Task struct {
	gorm.Model
	Details string
}
