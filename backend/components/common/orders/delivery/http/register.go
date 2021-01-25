package http

import (
	"github.com/abylq/learning-management-system/components/common/orders"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup,uc orders.UseCase) {
	h := NewHandler(uc)

	orders := router.Group("/orders")
	{
		orders.POST("",h.Create)
	}
}
