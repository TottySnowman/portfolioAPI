package authenticationModel

import "github.com/golang-jwt/jwt/v4"

type AuthenticationResponse struct {
	Token string
}

type JWTClaims struct {
	Username string
	jwt.RegisteredClaims
}
