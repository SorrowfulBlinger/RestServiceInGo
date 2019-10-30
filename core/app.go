package core

import "github.com/gin-gonic/gin"

func StartApplication() {
	router := gin.Engine{}
	router.POST("/login", nil)
	router.Run(":8080")
}
