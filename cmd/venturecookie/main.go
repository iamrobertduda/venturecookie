package main

import (
	internal "VentureCookie1/internal/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "inject")
	router.POST("/user/create", internal.PostUser)
	router.POST("/user/update", internal.UpdateUser)

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
