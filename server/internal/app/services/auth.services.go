package services

import (
	"context"

	"github.com/chihabMe/ichat/server/internal/app/models"
	"github.com/chihabMe/ichat/server/internal/app/repositories"
)


type AuthService struct {
    tokenRepository repositories.TokenRepository
}
func NewAuthService(tokenRepository repositories.TokenRepository)*AuthService{
    return &AuthService{tokenRepository: tokenRepository}
}
func (s *AuthService)SaveRefreshToken(ctx context.Context,token *models.Token)error{
	return s.tokenRepository.Create(ctx,token)
}

func (s *AuthService) DeleteRefreshTokenIfExisted(ctx context.Context,tokenString string)error{
    return s.tokenRepository.DeleteByTokenString(ctx,tokenString)
}


// func SaveRefreshToken(user *models.User,token string,exp int64)error {
// 	db := core.Instance
// 	 t:= models.Token{
// 		Token: token,
// 		UserID: user.ID,
// 		Exp: exp,

// 	}
// 	err := db.Create(&t).Error
// 	return err
// }
// func DeleteRefreshTokenIfExisted(t string) (*models.Token, error) {
//     db := core.Instance
//     var token models.Token
//     result := db.Where("token = ?", t).First(&token)
//     if result.Error != nil {
//         if errors.Is(result.Error, gorm.ErrRecordNotFound) {
//             return nil, errors.New("token does not exist")
//         }
//         return nil, result.Error
//     }

//     if err := db.Delete(&token).Error; err != nil {
//         return nil, errors.New("unable to delete the token")
//     }

//     return &token, nil
// }
