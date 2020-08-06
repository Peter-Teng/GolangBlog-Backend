package utils

import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type BlogClaims struct {
	IntId int64
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(username string, role string, id int64) (string, error) {
	claims := BlogClaims{
		id,
		role,
		jwt.StandardClaims{
			Audience:  username,
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
			Issuer:    "PP同学",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(Str2bytes(config.JwtSecurity))
	if err != nil {
		return "", err
	}
	token = "Bearer " + token
	return token, nil
}

func ParseToken(token string) (*BlogClaims, int) {
	var claims BlogClaims

	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return Str2bytes(config.JwtSecurity), nil
	})

	if err != nil {
		if problem, ok := err.(jwt.ValidationError); ok {
			if problem.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, common.TOKEN_WRONG_TOKEN
			} else if problem.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, common.TOKEN_NOT_VALID
			} else {
				return nil, common.TOKEN_WRONG_TYPE
			}
		}
	}

	if jwtToken != nil {
		if claims, ok := jwtToken.Claims.(*BlogClaims); ok && jwtToken.Valid {
			return claims, common.SUCCESS
		} else {
			return nil, common.TOKEN_WRONG_TOKEN
		}
	}
	return nil, common.TOKEN_WRONG_TOKEN
}
