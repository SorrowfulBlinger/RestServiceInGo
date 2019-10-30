package controller

import (
	"RestServiceInGo/dto"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

const (
	BadRequestFailure int8 = -1
	Success         int8 = 0
)

func HandleLogin(ctx *gin.Context) {
	request := dto.LoginRequestDto{}
	if err := ctx.ShouldBindWith(&request, binding.JSON);  err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.LoginResponseDto{StatusCode: BadRequestFailure})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.LoginResponseDto{StatusCode: Success})
}