package controller

import (
	"backend/middleware"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	//input := &LoginInput{}
	//fmt.Println(input.Username, input.Password)
	//if err := c.ShouldBindJSON(input); err != nil {
	//	utils.ErrorHandler(c, err)
	//	return
	//}

	username := c.Query("username")
	password := c.Query("password")

	//user, err := utils.ValidateLogin(input.Username, input.Password)
	user, err := utils.ValidateLogin(username, password)
	if err != nil {
		utils.ErrorHandler(c, err)
		return
	}

	token, err := middleware.GenerateToken(user.Username, user.ID)
	if err != nil {
		utils.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "login success",
		Data: gin.H{
			"user":  user,
			"token": token,
		},
	})
}
