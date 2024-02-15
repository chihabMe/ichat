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
	panic("unimplemented")
}

// Delete implements UserRepository.
func (r *UserRepositoryImpl) Delete(ctx context.Context, userId string) error {
	panic("unimplemented")
}

// FindByEmail implements UserRepository.
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, userEmail string) (*models.User, error) {
	panic("unimplemented")
}

// FindById implements UserRepository.
func (r *UserRepositoryImpl) FindById(ctx context.Context, userId string) (*models.User, error) {
	panic("unimplemented")
}

// FindByUsername implements UserRepository.
func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, userUsername string) (*models.User, error) {
	panic("unimplemented")
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
