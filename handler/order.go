package handler

import (
	"h8-assignment-2/dto"
	"h8-assignment-2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) orderHandler {
	return orderHandler{
		OrderService: orderService,
	}
}

func (oh *orderHandler) CreateOrder(ctx *gin.Context) {
	var newOrderRequest dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid json request",
		})
		return
	}

	err := oh.OrderService.CreateOrder(newOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully creating a product",
	})
}

func (oh *orderHandler) GetOrders(ctx *gin.Context) {
	response, err := oh.OrderService.GetOrders()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (oh *orderHandler) UpdateOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	var updateOrderRequest dto.UpdateOrderRequest

	if err := ctx.ShouldBindJSON(&updateOrderRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid json request",
		})
		return
	}

	err := oh.OrderService.UpdateOrder(orderId, updateOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully updating an order",
	})
}

func (oh *orderHandler) DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	err := oh.OrderService.DeleteOrder(orderId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully deleting an order",
	})
}
