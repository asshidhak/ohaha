package main

import (
	"github.com/gin-gonic/gin"
	"jintmp/api/host"
)

func initRoute() *gin.Engine{
	route := gin.Default()
	route.POST("/", host.HandleHostList)
	route.POST("/login")
	return route
}
