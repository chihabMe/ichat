package utils

import (
	"time"

	"github.com/chihabMe/ichat/server/core"
	"github.com/chihabMe/ichat/server/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.User)(string,error){
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"]=user.Username
	claims["user_id"] = user.ID
	claims["exp"]= time.Now().Add(core.ACCESS_TOKEN_TIME).Unix()

	secret_key := core.Config("SECRET_KEY")
	t,err := token.SignedString([]byte(secret_key))
	return t,err
}