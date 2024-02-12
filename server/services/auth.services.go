package services

import (
	"github.com/chihabMe/ichat/server/core"
	"github.com/chihabMe/ichat/server/models"
)

func SaveRefreshToken(user *models.User,token string,exp int64)error {
	db := core.Instance
	 t:= models.Token{
		Token: token,
		UserID: user.ID,
		Exp: exp,

	}
	err := db.Create(&t).Error
	return err
}