package httpserver

import (
	"assignment2/httpserver/controllers"

	"github.com/gin-gonic/gin"
)
type Router struct {
	router  *gin.Engine
	order *controllers.OrderController
}

func NewRouter(router *gin.Engine, order *controllers.OrderController) *Router {
	return &Router{
		router:  router,
		order: order,
	}
}

// func CreateRouter() *gin.Engine {
// 	router := gin.Default()

// 	router.POST("/orders", controllers.CreateOrder)
// 	router.GET("/orders", controllers.GetOrders)
// 	router.PUT("/orders/:id", controllers.UpdateOrder)
// 	router.DELETE("/orders/:id", controllers.DeleteOrder)

// 	return router
// }

func (r *Router) Start(port string) {
	// step :(1) request masuk, request keluar
	r.router.POST("/orders", r.order.CreateOrder)
	r.router.GET("/orders", r.order.GetAllOrders)
	r.router.PUT("/orders", r.order.UpdateOrder)
	r.router.DELETE("/orders/:id", r.order.DeleteOrder)
	r.router.Run(port)
}