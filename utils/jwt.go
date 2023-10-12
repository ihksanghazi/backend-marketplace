package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type customClaims struct{
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken (signingKey string,duration time.Time, id string, subject string, username string) (string,error) {
	claims := customClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ID: id,
			Subject: subject,
			ExpiresAt: jwt.NewNumericDate(duration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(signingKey))

	return ss,err
}

func ParsingToken(yourToken string,signingKey string) (*customClaims,error) {
	token, err := jwt.ParseWithClaims(yourToken, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		return claims,nil
	}else{
		return nil, err
	}

}