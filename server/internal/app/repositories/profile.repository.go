package repositories

import (
	"context"

	"github.com/chihabMe/ichat/server/internal/app/models"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(ctx context.Context, profile *models.Profile) error
	Update(ctx context.Context, profile *models.Profile) error
	DeleteByUserID(ctx context.Context, userID string) error
	All(ctx context.Context, profiles *[]models.Profile) error
}

type ProfileRepositoryImpl struct {
	db *gorm.DB
}


func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &ProfileRepositoryImpl{db: db}
}

func (r *ProfileRepositoryImpl) All(ctx context.Context, profiles *[]models.Profile) error {
	return r.db.WithContext(ctx).Model(&models.Profile{}).Find(profiles).Error
}

func (r *ProfileRepositoryImpl) Create(ctx context.Context, profile *models.Profile) error {
	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *ProfileRepositoryImpl) DeleteByUserID(ctx context.Context, userID string) error {
	return r.db.WithContext(ctx).Where("user_id = ?",userID).Delete(&models.Profile{}).Error
}

func (r *ProfileRepositoryImpl) Update(ctx context.Context, profile *models.Profile) error {
	return r.db.WithContext(ctx).Save(profile).Error
}
