package projectModel

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
  gorm.Model
	Name             string        `gorm:"size:255;not null"`           // Name of the project
	About            string        `gorm:"size:512;not null"`           // Description of the project
	GithubLink       string        `gorm:"size:255;not null"`           // GitHub repository link
	DemoLink         string        `gorm:"size:255;not null"`           // Demo or live site link
	LogoPath         string        `gorm:"size:255;not null"`           // Path to the project logo
	DevDate          time.Time     `gorm:"type:date;not null"`          // Development date
	Hidden           bool          `gorm:"not null"`                    // Hidden status (boolean)
  ProjectStatusID uint       `gorm:"not null"` // Foreign key field
  ProjectStatus   ProjectStatus `gorm:"foreignkey:ProjectStatusID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key constraint
	Tags             []Tag         `gorm:"many2many:Project_Tags"`      // Many-to-many relationship with Tag
}

// ProjectStatus struct corresponds to the `ProjectStatus` table
type ProjectStatus struct {
  gorm.Model
	Status string `gorm:"size:12;not null"`         // Status of the project
  Project []Project
}

// Tag struct corresponds to the `Tag` table
type Tag struct {
  gorm.Model
	Tag  string `gorm:"size:20"`                  // Name of the tag
	Icon string `gorm:"size:20;not null"`         // Icon representing the tag
}

type ProjectDataSelect struct{
  ProjectID int
  Status string
  Name string
  About string
  GithubLink string
  DemoLink string
  LogoPath string
  TagIcon string
  Tag string
}

type JsonTag struct {
  TagIcon string
  Tag string
}

type ProjectDisplay struct{
  ProjectID int
  Status string
  Name string
  About string
  Github_Link string
  Demo_Link string
  Logo_Path string
  Tags []JsonTag
}
