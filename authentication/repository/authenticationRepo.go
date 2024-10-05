package authenticationRepo

import (
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
