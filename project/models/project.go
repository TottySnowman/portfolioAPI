package projectModel

import(
  "time"
)

type Project struct {
	ID             int            `gorm:"primaryKey;autoIncrement"`   // Auto-increment primary key
	Name           string         `gorm:"size:255;not null"`          // Name of the project
	About          string         `gorm:"size:512;not null"`          // Description of the project
	GithubLink     string         `gorm:"size:255;not null"`          // GitHub repository link
	DemoLink       string         `gorm:"size:255;not null"`          // Demo or live site link
	LogoPath       string         `gorm:"size:255;not null"`          // Path to the project logo
	DevDate        time.Time      `gorm:"type:date;not null"`         // Development date
	Hidden         bool           `gorm:"not null"`                   // Hidden status (boolean)
	ProjectStatus  ProjectStatus  `gorm:"foreignKey:FK_Project_Status"` // Foreign key relation to ProjectStatus
	FK_ProjectStatus int          `gorm:"not null"`                   // Foreign key
	Tags           []Tag          `gorm:"many2many:Project_Tags"`     // Many-to-many relationship with Tag
}

// ProjectStatus struct corresponds to the `ProjectStatus` table
type ProjectStatus struct {
	ID     int    `gorm:"primaryKey;autoIncrement"` // Auto-increment primary key
	Status string `gorm:"size:12;not null"`         // Status of the project
}

// ProjectTags struct represents the `Project_Tags` table (for the many-to-many relationship between Project and Tag)
type ProjectTags struct {
	ProjectID int `gorm:"primaryKey"` // Composite primary key
	TagID     int `gorm:"primaryKey"` // Composite primary key
}

// Tag struct corresponds to the `Tag` table
type Tag struct {
	ID   int    `gorm:"primaryKey;autoIncrement"` // Auto-increment primary key
	Tag  string `gorm:"size:20"`                  // Name of the tag
	Icon string `gorm:"size:20;not null"`         // Icon representing the tag
}
