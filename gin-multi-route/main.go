package main

import (
	"gin-multi-route/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.LoginRoutes(router)
	routes.RegisterRoutes(router)

	router.Run()
}
