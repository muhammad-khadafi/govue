package handler

/*
 * Created by muhammad.khadafi on 27/05/2022
 */

import (
	"backend/request"
	"backend/service"
	"backend/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"html"
	"net/http"
	"strconv"
	"strings"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (userHandler *userHandler) HelloWithoutAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Hello World",
	})
}

func (userHandler *userHandler) FindUserByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	user, err := userHandler.userService.FindUserByID(uint(id))
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Retrieve data success!", http.StatusOK, "ok", user)
}

func (userHandler *userHandler) InsertUser(c *gin.Context) {
	var user request.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	//turn password into hash
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errHash != nil {
		util.APIResponse(c, errHash.Error(), http.StatusBadRequest, "error", nil)
	}
	user.Password = string(hashedPassword)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	cekEmail, err := userHandler.userService.FindUserByEmail(user.Email)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
			return
		}
	}
	if cekEmail.Id != 0 {
		util.APIResponse(c, "Email telah terdaftar", http.StatusBadRequest, "error", nil)
		return
	}
	cekUsername, err := userHandler.userService.FindUserByUsername(user.Username)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
			return
		}
	}
	if cekUsername.Id != 0 {
		util.APIResponse(c, "Username telah terdaftar", http.StatusBadRequest, "error", nil)
		return
	}

	err = userHandler.userService.InsertUser(user)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Insert data success!", http.StatusOK, "ok", nil)
}
