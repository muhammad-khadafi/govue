package util

import (
	"backend/entity"
	"backend/repository"
	"backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
	"strings"
	"time"
)

func GenerateToken(user_id uint) (string, error) {

	token_lifespan, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_MINUTE_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

func GenerateRefreshToken(user_id uint) (string, error) {
	refresh_token_lifespan, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(refresh_token_lifespan)).Unix()
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refresh_token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func RefreshTokenValid(c *gin.Context) error {
	tokenString := ExtractRefreshToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractRefreshToken(c *gin.Context) string {
	refreshToken := c.Request.Header.Get("refresh_token")
	if refreshToken != "" {
		return refreshToken
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

func ExtractRefreshTokenToMap(c *gin.Context) (entity.TokenEntity, error) {
	var tokenEnt entity.TokenEntity
	tokenString := ExtractRefreshToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return tokenEnt, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			uid = 0
		}
		exp, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["exp"]), 10, 32)
		if err != nil {
			exp = 0
		}
		expTime := time.Unix(exp, 0)
		tokenEnt.UserId = uint(uid)
		tokenEnt.ExpiredTime = expTime

		return tokenEnt, nil
	}
	return tokenEnt, nil
}

func InsertToBlacklist(c *gin.Context, tokenEntity entity.TokenEntity, db *sqlx.DB) {
	tokenRepository := repository.NewTokenRepository(db)
	tokenService := service.NewTokenService(tokenRepository)

	var blacklistToken entity.BlacklistedRefreshToken
	blacklistToken.Token = c.Request.Header.Get("refresh_token")
	blacklistToken.CreatedAt = time.Now()
	blacklistToken.ExpiredAt = tokenEntity.ExpiredTime

	tokenService.InsertToken(blacklistToken)
}

func CheckRefreshToken(token string, db *sqlx.DB) string {
	tokenRepository := repository.NewTokenRepository(db)
	tokenService := service.NewTokenService(tokenRepository)

	blacklistedToken, _ := tokenService.FindByToken(token)
	if blacklistedToken.Id != 0 {
		return "The token has been blacklisted, please re-login"
	}
	return ""
}
