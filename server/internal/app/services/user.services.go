package services

import (
	"context"

	"github.com/chihabMe/ichat/server/internal/app/models"
	"github.com/chihabMe/ichat/server/internal/app/repositories"
)


type UserService struct {
	userRepository repositories.UserRepository
	profileRepository repositories.ProfileRepository
}

func NewUserService(userRepository repositories.UserRepository,profileRepository repositories.ProfileRepository )*UserService{
	return &UserService{userRepository: userRepository,profileRepository: profileRepository}

}

func (s *UserService)CreateUser(ctx context.Context,user *models.User,profile *models.Profile)error{
	if err:= s.userRepository.Create(ctx,user);err!=nil{
		return err
	}
	profile.UserId = user.ID
	if err:= s.profileRepository.Create(ctx,profile);err!=nil{
		return err
	}
	return nil
}
func (s *UserService)UpdateUser(ctx context.Context,user *models.User)error{
	return s.userRepository.Update(ctx,user)
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
    return s.userRepository.Delete(ctx, userID)
}

func (s *UserService) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
    return s.userRepository.FindByID(ctx, userID)
}

func (s *UserService) GetUserWithProfileByID(ctx context.Context, userID string) (*models.User, error) {
    return s.userRepository.FindByIDWithProfile(ctx, userID)
}
func (s *UserService) GetAllUsers(ctx context.Context,users *[]models.User)error{
	return s.userRepository.All(ctx,users)
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
    return s.userRepository.FindByUsername(ctx, username)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
    return s.userRepository.FindByEmail(ctx, email)
}


func (s *UserService) UpdateUserPassword(ctx context.Context,user *models.User,newPassword string)  error {
	user.Password = newPassword
	return  s.userRepository.Update(ctx,user)
}

