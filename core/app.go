package core

import (
	"RestServiceInGo/controller"
	"github.com/gin-gonic/gin"
)

func StartApplication() {
	router := gin.Default()

	router.POST("/login", controller.HandleLogin)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}
