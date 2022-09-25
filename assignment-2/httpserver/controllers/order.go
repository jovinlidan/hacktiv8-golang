package controllers

import (
	"assignment2/httpserver/controllers/params"
	"assignment2/httpserver/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderController struct {
	svc *services.OrderSvc
}

func NewOrderController(svc *services.OrderSvc) *OrderController {
	return &OrderController{
		svc: svc,
	}
}

func (s *OrderController) GetAllOrders(ctx *gin.Context) {
	response := s.svc.GetAllOrders()
	WriteJsonRespnse(ctx, response)
}

func (s *OrderController) UpdateOrder(ctx *gin.Context) {
	var req params.OrderUpdateRequest
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

	response := s.svc.UpdateOrder(&req)
	WriteJsonRespnse(ctx, response)
}

func (s *OrderController) DeleteOrder(ctx *gin.Context) {
	req := ctx.Param("id")
	

	response := s.svc.DeleteOrder(&req)
	WriteJsonRespnse(ctx, response)
}

func (s *OrderController) CreateOrder(ctx *gin.Context) {
	var req params.OrderCreateRequest
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

	response := s.svc.CreateOrder(&req)
	WriteJsonRespnse(ctx, response)

}
