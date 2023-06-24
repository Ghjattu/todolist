package main

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	apiRouter := r.Group("/todolist/v1")
	// user module
	apiRouter.POST("/user/register", controllers.Register)
	apiRouter.POST("/user/login", controllers.Login)
	apiRouter.POST("/user/logout")

	// item module
	apiRouter.GET("/item")
	apiRouter.POST("/item/new")
	apiRouter.PUT("/item/:id")
	apiRouter.DELETE("/item/:id")
}
