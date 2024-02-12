package services

import (
	"errors"

	"github.com/chihabMe/ichat/server/core"
	"github.com/chihabMe/ichat/server/models"
	"gorm.io/gorm"
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
func DeleteRefreshTokenIfExisted(t string) (*models.Token, error) {
    db := core.Instance
    var token models.Token
    result := db.Where("token = ?", t).First(&token)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("token does not exist")
        }
        return nil, result.Error
    }

    if err := db.Delete(&token).Error; err != nil {
        return nil, errors.New("unable to delete the token")
    }

    return &token, nil
}
