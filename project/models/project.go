package projectModel

import (
	"gorm.io/gorm"
	statusModel "portfolioAPI/status/models"
	tagModel "portfolioAPI/tag/models"
	"time"
)

type Project struct {
	gorm.Model
	ID              int
	Name            string         `gorm:"size:255;not null"`                                                         // Name of the project
	About           string         `gorm:"size:512;not null"`                                                         // Description of the project
	GithubLink      string         `gorm:"size:255;not null"`                                                         // GitHub repository link
	DemoLink        string         `gorm:"size:255;not null"`                                                         // Demo or live site link
	LogoPath        string         `gorm:"size:255;not null"`                                                         // Path to the project logo
	DevDate         time.Time      `gorm:"type:date;not null"`                                                        // Development date
	Hidden          bool           `gorm:"not null"`                                                                  // Hidden status (boolean)
	ProjectStatusID uint           `gorm:"not null"`                                                                  // Foreign key field
	ProjectStatus   ProjectStatus  `gorm:"foreignkey:ProjectStatusID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key constraint
	Tags            []tagModel.Tag `gorm:"many2many:Project_Tags"`                                                    // Many to many
	LanguageCode    string         `gorm:"not null;default:de-CH"`
}

type ProjectStatus struct {
	gorm.Model
	Status  string `gorm:"size:12;not null"`
	Project []Project
}

type ProjectDataSelect struct {
	ProjectID  int
	Status     string
	StatusID   int
	Name       string
	About      string
	GithubLink string
	DemoLink   string
	LogoPath   string
	TagIcon    string
	Tag        string
	TagId      int
	DevDate    time.Time
	Hidden     bool
	LanguageCode     string
}

type ProjectDisplay struct {
	ProjectID   int
	Status      statusModel.ProjectStatusDisplay
	Name        string
	About       string
	Github_Link string
	Demo_Link   string
	Logo_Path   string
	Tags        []tagModel.JsonTag
	DevDate     time.Time
	Hidden      bool
	LanguageCode      string
}
