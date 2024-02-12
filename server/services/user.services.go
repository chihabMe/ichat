package services

import (
	"errors"

	"github.com/chihabMe/ichat/server/core"
	"github.com/chihabMe/ichat/server/models"
	"gorm.io/gorm"
)


func GetUserByEmail(email string)(*models.User,error){
	db := core.Instance
	var user models.User
	if err:= db.Where(&models.User{Email: email}).Find(&user).Error;err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,nil
		}
		return nil,err
	}
	return &user,nil
}

func GetUserByID(id uint)(*models.User,error){
	db := core.Instance
	var user models.User
	if err:= db.Where("id = ?",id).Find(&user).Error;err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,nil
		}
		return nil,err
	}
	return &user,nil
}