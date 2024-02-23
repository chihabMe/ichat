package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/chihabMe/ichat/server/internal/app/config"
	"github.com/chihabMe/ichat/server/internal/app/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtClaims struct {
	Username string `json:"username"`
	UserId   uuid.UUID   `json:"user_id"`
	Exp      int64  `json:"exp"`
	jwt.RegisteredClaims
}
type TokenData struct {
	Body string
	Exp int64
}

func GenerateAccessToken(user *models.User)(TokenData,error){
	cfg := config.InitConfig()
	return generateToken(user,cfg.AccessTokenTTL)
}
func GenerateRefreshToken(user *models.User)(TokenData,error){
	cfg := config.InitConfig()
	return generateToken(user,cfg.RefreshTokenTTL)
}
func generateToken(user *models.User,expires time.Duration ) (TokenData, error) {
	c := JwtClaims{
		Username: user.Username,
		UserId:   user.ID,
		Exp:      time.Now().Add(expires).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)


	var secretKey = config.GetEnvOrDefault("SECRET_KEY","")

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


func VerifyToken(tokenString string)(*jwt.Token,error){
	token,err :=jwt.Parse(tokenString,func(token *jwt.Token)(interface{},error){
		if _,ok :=token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,fmt.Errorf("unexpected siding method ")
		}
		secretKey := config.GetEnvOrDefault("SECRET_KEY","")
		return []byte(secretKey),nil
	})
	if err!=nil{
		return nil,err
	}
	_,ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid{
		return nil,errors.New("Invalid token")
	}
	return token,nil

}
func VerifyAccessToken(tokenString string)(*jwt.Token,error){
	token,err := VerifyToken(tokenString)
	if err !=nil{
		return nil,err
	}
	alive :=VerifyTokenExpireDate(token)
	if !alive {
		return nil,errors.New("dead token")
	}
	return token,nil
}
func VerifyTokenExpireDate(token *jwt.Token) bool {
	claims := token.Claims.(jwt.MapClaims)
	exp := float64(claims["exp"].(float64))
	now := float64(time.Now().Unix())
	return now < exp
}