package authenticationModel

type User struct {
	Username string `gorm:"size:50;not null"` 
	Email    string `gorm:"size:100;not null"`
	Password string `gorm:"size:255;not null"`
}
