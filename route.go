package main

import (
	"github.com/gin-gonic/gin"
	"github.com/asshidhak/ohaha/api/host"
)

func initRoute() *gin.Engine{
	route := gin.Default()
	route.POST("/", host.HandleHostList)
	route.POST("/login")
	return route
}
