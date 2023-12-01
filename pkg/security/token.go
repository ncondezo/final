package security

import (
	"os"
	"time"

	"github.com/ncondezo/final/internal/domain"

	"github.com/golang-jwt/jwt"
)

var secretKey = os.Getenv("TOKEN_SECRET_KEY")

func GenerateToken(dto *domain.LoginDTO) (string, error) {
	claims := &domain.Claim{
		dto.Email,
		jwt.StandardClaims{
			Issuer:    "desafio2-backend",
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(token string) (*domain.Claim, error) {
	claim := &domain.Claim{}
	tkn, err := jwt.ParseWithClaims(token, claim,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
	if err != nil {
		return nil, err
	}
	if tkn.Valid {
		return claim, nil
	}
	return nil, jwt.NewValidationError("Invalid token.", 0)
}
