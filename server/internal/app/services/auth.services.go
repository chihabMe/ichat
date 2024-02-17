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

