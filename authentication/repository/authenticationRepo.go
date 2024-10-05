package authenticationRepo

import (
	authenticationModel "portfolioAPI/authentication/models"
	"portfolioAPI/database"

	"gorm.io/gorm"
)

type AuthRepo struct{
  db *gorm.DB
}

func NewAuthRepo() *AuthRepo{
  return &AuthRepo{
    db: database.GetDBClient(),
  }
}

func(repo *AuthRepo) AuthenticateUser(userInput authenticationModel.LoginRequest) *authenticationModel.User{
  return &authenticationModel.User{
    Username: "Bingbong",
  }
}
