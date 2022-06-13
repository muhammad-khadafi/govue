package repository

import (
	"backend/entity"
	"github.com/jmoiron/sqlx"
	"time"
)

type TokenRepository interface {
	Insert(token entity.BlacklistedRefreshToken) error
	DeleteExpiredToken() int
	FindByToken(token string) (entity.BlacklistedRefreshToken, error)
}

type tokenRepository struct {
	DB *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *tokenRepository {
	return &tokenRepository{DB: db}
}

func (repository *tokenRepository) Insert(token entity.BlacklistedRefreshToken) error {
	token.CreatedAt = time.Now()
	script := `INSERT INTO blacklisted_refresh_token (token, created_at, expired_at) VALUES (:TOKEN, :CREATED_AT, :EXPIRED_AT)`
	_, err := repository.DB.NamedExec(script, &token)
	return err
}

func (repository *tokenRepository) DeleteExpiredToken() int {
	script := `DELETE FROM blacklisted_refresh_token WHERE expired_at < :1`
	del := repository.DB.MustExec(script, time.Now())
	count, _ := del.RowsAffected()
	return int(count)
}

func (repository *tokenRepository) FindByToken(token string) (entity.BlacklistedRefreshToken, error) {
	blacklistedToken := entity.BlacklistedRefreshToken{}
	script := `select * from blacklisted_refresh_token where token = :1`
	err := repository.DB.Get(&blacklistedToken, script, token)
	return blacklistedToken, err
}
