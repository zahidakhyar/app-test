package auth_service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtServiceInterface interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type JwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() JwtServiceInterface {
	return &JwtService{
		secretKey: getSecretKey(),
		issuer:    "zahidakhyar",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	if secretKey == "" {
		secretKey = "secret"
	}

	return secretKey
}

func (s *JwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}

	return t
}

func (s *JwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t_.Header["alg"])
		}

		return []byte(s.secretKey), nil
	})
}
