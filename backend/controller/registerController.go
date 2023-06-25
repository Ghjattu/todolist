package controller

import (
	"backend/dao"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	//input := &RegisterInput{}
	//if err := c.ShouldBindJSON(input); err != nil {
	//	utils.ErrorHandler(c, err)
	//	return
	//}

	username := c.Query("username")
	password := c.Query("password")
	fmt.Println(username, password)
	user := &dao.User{
		Username: username,
		Password: password,
	}
	returnedUser, err := dao.CreateNewUser(user)
	if err != nil {
		utils.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "registration success",
		Data:    returnedUser,
	})
}
