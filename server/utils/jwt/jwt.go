package utils

import (
	"time"

	"github.com/chihabMe/ichat/server/core"
	"github.com/chihabMe/ichat/server/models"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Username string `json:"username"`
	UserId   uint   `json:"user_id"`
	Exp      int64  `json:"exp"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User) (string, error) {
	c := JwtClaims{
		Username: user.Username,
		UserId:   user.ID,
		Exp:      time.Now().Add(core.ACCESS_TOKEN_TIME).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)


	var secretKey = core.Config("SECRET_KEY")

	// Sign the token with the ECDSA private key
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
