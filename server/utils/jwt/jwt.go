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
type TokenData struct {
	Body string
	Exp int64
}

func GenerateAccessToken(user *models.User)(TokenData,error){
	return generateToken(user,core.ACCESS_TOKEN_TIME)
}
func GenerateRefreshToken(user *models.User)(TokenData,error){
	return generateToken(user,core.REFRESH_TOKEN_TIME)
}
func generateToken(user *models.User,expires time.Duration ) (TokenData, error) {
	c := JwtClaims{
		Username: user.Username,
		UserId:   user.ID,
		Exp:      time.Now().Add(expires).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)


	var secretKey = core.Config("SECRET_KEY")

	// Sign the token with the ECDSA private key
	t, err := token.SignedString([]byte(secretKey))
	 tokenData := TokenData{
		Body: t,
		Exp:c.Exp ,
	 }
	if err != nil {
		return tokenData, err
	}

	return tokenData, nil
}
