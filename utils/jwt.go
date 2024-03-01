package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(GetEnvKey("JWT_SECRET")))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return []byte(GetEnvKey("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, errors.New("token inválido")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("não foi possível parsear os claims")
	}

	if time.Unix(int64(claims["exp"].(float64)), 0).Before(time.Now()) {
		return 0, errors.New("token expirado")
	}

	// email := claims["email"]
	userId := int64(claims["userId"].(float64)) // o valor é float64, então precisamos converter para int64
	return userId, nil
}
