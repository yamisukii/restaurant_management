package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/yamisukii/restaurant_managementGo/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemdByOrder())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
}
