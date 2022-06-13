package response

import "backend/entity"

/*
 * Created by muhammad.khadafi on 27/05/2022
 */

type UserSessionResponse struct {
	Id           uint   `json:"id"`
	Username     string `json:"username"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func ToUserSessionResponse(user entity.User, token string, refreshToken string) UserSessionResponse {
	userSessionResponse := UserSessionResponse{
		Id:           user.Id,
		Username:     user.Username,
		Token:        token,
		RefreshToken: refreshToken,
	}

	return userSessionResponse
}
