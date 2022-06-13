package service

import (
	"backend/entity"
	"backend/repository"
)

type TokenService interface {
	InsertToken(token entity.BlacklistedRefreshToken) error
	DeleteExpiredToken() int
	FindByToken(token string) (entity.BlacklistedRefreshToken, error)
}

type tokenService struct {
	tokenRepository repository.TokenRepository
}

func NewTokenService(tokenRepository repository.TokenRepository) *tokenService {
	return &tokenService{tokenRepository}
}

func (service *tokenService) InsertToken(token entity.BlacklistedRefreshToken) error {
	err := service.tokenRepository.Insert(token)
	return err
}

func (service *tokenService) DeleteExpiredToken() int {
	err := service.tokenRepository.DeleteExpiredToken()
	return err
}

func (service *tokenService) FindByToken(token string) (entity.BlacklistedRefreshToken, error) {
	blacklistedToken, err := service.tokenRepository.FindByToken(token)
	return blacklistedToken, err
}
