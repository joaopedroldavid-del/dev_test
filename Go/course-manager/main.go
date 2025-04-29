package main

import (
	"github.com/gin-gonic/gin"
	"course-manager/config"
	"course-manager/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	routes.SetupRoutes(r)
	r.Run(":8080")
}