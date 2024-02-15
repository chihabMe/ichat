package services

import (
	"errors"

	"github.com/chihabMe/ichat/server/core"
	"github.com/chihabMe/ichat/server/models"
	"gorm.io/gorm"
)

func CreateUser(user *models.User)error{
	db:= core.Instance
	return db.Save(user).Error

}
func CreateProfile(profile *models.Profile)error{
	db:= core.Instance
	return db.Save(profile).Error
}
func UpdateUser(user *models.User)error{
	db := core.Instance
	return db.Save(user).Error
}

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
func UpdateUserPassword(newPasswordHash string,user *models.User)error{
	db:=core.Instance
	user.Password = newPasswordHash
	return db.Save(&user).Error

}