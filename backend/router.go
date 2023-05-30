package main

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	apiRouter := r.Group("/todolist/v1")
	// user module
	apiRouter.POST("/user/register")
	apiRouter.POST("/user/login")
	apiRouter.POST("/user/logout")

	// item module
	apiRouter.GET("/item")
	apiRouter.POST("/item/new")
	apiRouter.PUT("/item/:id")
	apiRouter.DELETE("/item/:id")
}
