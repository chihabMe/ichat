package repositories

import (
	"context"

	"github.com/chihabMe/ichat/server/internal/app/models"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Create(ctx context.Context, token *models.Token) error
	Delete(ctx context.Context, token *models.Token) error
	DeleteAllByUserID(ctx context.Context, userID string) error
	FindByTokenString(ctx context.Context, tokenString string) (*models.Token, error)
}

type TokenRepositoryImpl struct {
	db *gorm.DB
}
func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &TokenRepositoryImpl{db: db}
}

func (r *TokenRepositoryImpl) Create(ctx context.Context, token *models.Token) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *TokenRepositoryImpl) Delete(ctx context.Context, token *models.Token) error {
	return r.db.WithContext(ctx).Delete(token).Error
}

// FindByTokenString implements TokenRepository.
func (r *TokenRepositoryImpl) FindByTokenString(ctx context.Context, tokenString string) (*models.Token, error) {
	var token models.Token
	if err:= r.db.WithContext(ctx).Where("token = ?",tokenString).First(&token).Error; err!=nil{
		return nil,err
	}
	return &token,nil
}

func (r *TokenRepositoryImpl) DeleteAllByUserID(ctx context.Context,userID string)error{
	return r.db.WithContext(ctx).Where("user_id = ?",userID).Delete(&models.Token{}).Error

}