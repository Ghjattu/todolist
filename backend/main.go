package main

import (
	"backend/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})
	InitRouter(r)
	initDependencies()
	err := r.Run()
	if err != nil {
		return
	}
}

func initDependencies() {
	dao.Init()
}
