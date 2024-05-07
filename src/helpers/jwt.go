package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

type (
	ParamsGenerateJWT struct {
		ExpiredInMinute int
		SecretKey       string
		UserId          int64
		Role            string
	}
	ParamsValidateJWT struct {
		Token     string
		SecretKey string
	}
	Claims struct {
		jwt.StandardClaims
		UserId int64  `json:"user_id,omitempty"`
		Role   string `json:"role,omitempty"`
	}
)

func GenerateJWT(p *ParamsGenerateJWT) (string, int64, error) {
	expiredAt := time.Now().Add(time.Duration(p.ExpiredInMinute) * time.Minute).Unix()
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
		UserId: p.UserId,
		Role:   p.Role,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString([]byte(p.SecretKey))

	return signedToken, expiredAt, err
}

func ValidateJWT(p *ParamsValidateJWT) (jwt.MapClaims, error) {
	token, err := jwt.Parse(p.Token, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != JWT_SIGNING_METHOD {
			return nil, errors.New("token tidak valid")
		}

		return []byte(p.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
