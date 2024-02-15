package repositories

import (
	"context"

	"github.com/chihabMe/ichat/server/internal/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userId string) error
	FindById(ctx context.Context, userId string) (*models.User, error)
	FindByEmail(ctx context.Context, userEmail string) (*models.User, error)
	FindByUsername(ctx context.Context, userUsername string) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// Create implements UserRepository.
func (r *UserRepositoryImpl) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
func (r *UserRepositoryImpl) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Model(user).Updates(user).Error
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, userId string) error {
	return r.db.WithContext(ctx).Delete(&models.User{},userId).Error
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, userEmail string) (*models.User, error) {
	var user models.User
	if err:=r.db.WithContext(ctx).Where("email = ?",userEmail).First(&user).Error;err!=nil{
		return nil,err
	}
	return &user,nil
}

func (r *UserRepositoryImpl) FindById(ctx context.Context, userId string) (*models.User, error) {

	var user models.User
	if err:=r.db.WithContext(ctx).First(&user,userId).Error;err!=nil{
		return nil,err
	}
	return &user,nil
}

func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, userUsername string) (*models.User, error) {
	var user models.User
	if err:= r.db.WithContext(ctx).Where("username = ?",userUsername).First(&user).Error;err!=nil{
		return nil,err
	}
	return &user,nil
}


func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
