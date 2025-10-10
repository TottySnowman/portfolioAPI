package journeyModels

import (
	"gorm.io/gorm"
	"time"
)

type ExperienceType struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Experiences []Experience
}

type Experience struct {
	gorm.Model
  ID               int
	Title            string     `gorm:"not null"`
	Company          string     `gorm:"not null"`
	From             time.Time  `gorm:"not null"`
	To               *time.Time `json:"to,omitempty"`
	Diploma          string     `json:"diploma,omitempty"`
  LanguageCode     string     `gorm:"not null;default:de-CH"`
	Tasks            []Task     `json:"tasks,omitempty"`
	ExperienceTypeId uint
	ExperienceType   ExperienceType
}

type ExperienceModifyModel struct {
	Title          string             `json:"title"`
	Diploma        string             `json:"diploma"`
	From           time.Time          `json:"from"`
	To             *time.Time         `json:"to"`
	Tasks          []TaskDisplay      `json:"tasks"`
	ExperienceType ExperienceTypeEnum `json:"experienceType"`
}

type Task struct {
	gorm.Model
	Details      string `gorm:"not null"`
	ExperienceId uint
}

type ExperienceSelect struct {
	ExperienceId     uint
	Title            string
	Company          string
	From             time.Time
	To               *time.Time
	Name             string
	Diploma          string
	Details          string
	ExperienceTypeId uint
}

type ExperienceTypeEnum int

const (
	Education ExperienceTypeEnum = iota + 1
	Work
)

type JourneyDisplay struct {
	Id             int                `json:"id"`
	Title          string             `json:"title"`
	Diploma        string             `json:"diploma"`
	From           time.Time          `json:"from"`
	To             *time.Time         `json:"to"`
	Tasks          []TaskDisplay      `json:"tasks"`
	ExperienceType ExperienceTypeEnum `json:"experienceType"`
}

type TaskDisplay struct {
	Details string
}

type JourneyResponse struct {
	Educations []JourneyDisplay `json:"educations"`
	Work       []JourneyDisplay `json:"work"`
}
