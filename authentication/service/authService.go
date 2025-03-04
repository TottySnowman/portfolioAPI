package authService

import (
	"log"
	"os"
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

func (service *AuthService) AuthenticateUser(userInput *authenticationModel.LoginRequest) *authenticationModel.AuthenticationResponse {
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
  jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &authenticationModel.AuthenticationResponse{
		Token: tokenString,
    Username: authenticatedUser.Username,
	}
}

func (service *AuthService) IsAdminCreated() bool{
  return service.repo.IsAdminCreated()
}

func(service *AuthService) RegisterAdmin() bool{
  return service.repo.RegisterAdmin()
}
