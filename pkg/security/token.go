package security

import (
	"time"

	"github.com/ncondezo/final/internal/domain"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(dto *domain.LoginDTO) (string, error) {
	claim := &domain.Claim{
		dto.Email,
		jwt.StandardClaims{
			Issuer:    "desafio2-backend",
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	return token.SignedString(privateKey)
}

func ValidateToken(token string) (*domain.Claim, error) {
	claim := &domain.Claim{}
	tkn, err := jwt.ParseWithClaims(token, claim,
		func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})
	if err != nil {
		return nil, err
	}
	if tkn.Valid {
		return claim, nil
	}
	return nil, jwt.NewValidationError("Invalid token.", 0)
}
