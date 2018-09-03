package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	route := initRoute()
	route.Run()
}