package http

import (
	"github.com/abylq/learning-management-system/components/admin/quiz"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Quiz struct {
	ID 		 	int64 `json:"id"`
	ParentID 	int64 `json:"parent_id"`
	Title 	 	string `json:"title"`
	CreatedAT  	string
	UpdatedAT  	string
}

type Handler struct {
	useCase quiz.UseCase
}

func NewHandler(useCase quiz.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInput struct {
	Title string `json:"title"`
}

func (h *Handler) CreateQuizCategory(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreateCategoryQuiz(c.Request.Context(), inp.Title); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}