package controllers

import (
	"clean-arch-1/httpserver/controllers/params"
	"clean-arch-1/httpserver/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUser(ctx *gin.Context) {
	var req params.UserCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := services.CreateUser(&req)

	WriteJsonRespnse(ctx, response)
}

func GetUsers(ctx *gin.Context) {
	response := services.GetUsers()

	WriteJsonRespnse(ctx, response)
}
