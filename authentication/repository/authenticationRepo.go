package authenticationRepo

import (
	"os"
	authenticationModel "portfolioAPI/authentication/models"
	"portfolioAPI/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{
		db: database.GetDBClient(),
	}
}

func (repo *AuthRepo) AuthenticateUser(userInput *authenticationModel.LoginRequest) *authenticationModel.User {
	var existingUser *authenticationModel.User

	result := repo.db.Where("email = ?", userInput.Username).First(&existingUser)
	if result.Error != nil || existingUser.Username == "" {
		return nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userInput.Password))
	if err != nil {
		return nil
	}

	return &authenticationModel.User{
		Username: existingUser.Username,
	}
}

func (repo *AuthRepo) IsAdminCreated() bool {
	var count int64
	result := repo.db.Table("User").Count(&count)

	if result.Error != nil || count == 0 {
		return false
	}
	return true
}

func (repo *AuthRepo) RegisterAdmin() bool {
	if repo.IsAdminCreated() {
		return false
	}

	newAdminUser := &authenticationModel.User{
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		Username: os.Getenv("ADMIN_USERNAME"),
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newAdminUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}

  newAdminUser.Password = string(hashedPassword)

	result := repo.db.Create(&newAdminUser)
	if result.Error != nil {
		return false
	}

	return true
}
