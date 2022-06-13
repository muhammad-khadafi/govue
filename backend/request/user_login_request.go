package request

/*
 * Created by muhammad.khadafi on 27/05/2022
 */

type UserLoginRequest struct {
	//terima email/username
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
