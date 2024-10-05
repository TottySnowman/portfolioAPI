package authService

import (
	"log"
	authenticationModel "portfolioAPI/authentication/models"
	authenticationRepo "portfolioAPI/authentication/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	repo authenticationRepo.AuthRepo
}

func NewAuthService() *AuthService {
	return &AuthService{
		repo: *authenticationRepo.NewAuthRepo(),
	}
}

func (service *AuthService) AuthenticateUser(userInput authenticationModel.LoginRequest) *authenticationModel.AuthenticationResponse {
	authenticatedUser := service.repo.AuthenticateUser(userInput)
	if authenticatedUser == nil {
		return nil
	}

  expirationTime := time.Now().Add(24 * time.Hour)
	claims := &authenticationModel.JWTClaims{
		Username: authenticatedUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime)},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  mySigningKey := []byte("AllYourBase")
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &authenticationModel.AuthenticationResponse{
		Token: tokenString,
	}
}
