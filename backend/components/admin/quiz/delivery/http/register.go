package http

import (
	"github.com/abylq/learning-management-system/components/admin/quiz"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup,uc quiz.UseCase) {
	h := NewHandler(uc)

	quiz := router.Group("/quiz")
	{
		quiz.POST("category",h.CreateQuizCategory)
	}
}
