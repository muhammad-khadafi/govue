package middlewares

import (
	"backend/util"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func JwtAuthMiddleware(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := util.TokenValid(c)
		if err != nil {
			if err.Error() != "Token is expired" {
				c.String(http.StatusUnauthorized, err.Error())
				c.Abort()
				return
			}
			err := util.RefreshTokenValid(c)
			if err != nil {
				c.String(http.StatusUnauthorized, err.Error())
				c.Abort()
				return
			}
			refreshToken := c.Request.Header.Get("refresh_token")
			errCheck := util.CheckRefreshToken(refreshToken, db)

			if errCheck != "" {
				c.String(http.StatusUnauthorized, errCheck)
				c.Abort()
				return
			}
			tokenEnt, _ := util.ExtractRefreshTokenToMap(c)
			util.InsertToBlacklist(c, tokenEnt, db)
			token, err := util.GenerateToken(tokenEnt.UserId)
			if err != nil {
				c.String(http.StatusUnauthorized, err.Error())
				c.Abort()
				return
			}
			newRefreshToken, err := util.GenerateRefreshToken(tokenEnt.UserId)
			if err != nil {
				c.String(http.StatusUnauthorized, err.Error())
				c.Abort()
				return
			}
			c.Set("newToken", map[string]interface{}{
				"token":         token,
				"refresh_token": newRefreshToken,
			})
		}
		c.Next()
	}
}
