package handler

import (
	"backend/entity"
	"backend/request"
	"backend/response"
	"backend/service"
	"backend/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type loginHandler struct {
	userService service.UserService
}

func NewLoginHandler(userService service.UserService) *loginHandler {
	return &loginHandler{userService}
}

func (loginHandler *loginHandler) Login(c *gin.Context) {

	var userRequest request.UserLoginRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	token, refresh_token, user, err := loginHandler.loginCheck(userRequest.Email)
	if err != nil {
		util.APIResponse(c, "username/email or password is incorrect.", http.StatusBadRequest, "error", nil)
		return
	}

	userSessionResponse := response.ToUserSessionResponse(user, token, refresh_token)
	util.APIResponse(c, "Login success!", http.StatusOK, "ok", userSessionResponse)
}

func (loginHandler *loginHandler) loginCheck(input string) (string, string, entity.User, error) {
	user, err := loginHandler.userService.FindUserByUsernameOrEmail(input)
	if err != nil {
		return "", "", user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", "", user, err
	}

	token, err := util.GenerateToken(user.Id)
	if err != nil {
		return "", "", user, err
	}

	refresh_token, err := util.GenerateRefreshToken(user.Id)
	if err != nil {
		return "", "", user, err
	}

	return token, refresh_token, user, nil
}

func (loginHandler *loginHandler) CurrentUser(c *gin.Context) {
	user_id, err := util.ExtractTokenID(c)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	u, err := loginHandler.userService.FindUserByID(user_id)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Successfully retrieve data", http.StatusOK, "ok", u)
}

func (loginHandler *loginHandler) Logout(c *gin.Context) {
	util.APIResponseLogout(c, "Logout success!", http.StatusOK, "ok", nil)
}
