package http

import (
	"github.com/abylq/learning-management-system/components/auth"
	"github.com/abylq/learning-management-system/components/auth/models"
	"github.com/abylq/learning-management-system/components/common/orders"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Order struct {
	ID int64 `json:"id"`
	UserID int `json:"user_id"`
	DateStart string `json:"date_start"`
	Items int `json:"items"`
	Discipline string `json:"discipline"`
	Status int `json:"status"`
}

type Handler struct {
	useCase orders.UseCase
}

func NewHandler(useCase orders.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInput struct {
	DateStart string `json:"date_start"`
	Items int `json:"items"`
	Discipline int64 `json:"discipline"`
	Status int `json:"status"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreateOrder(c.Request.Context(), user,
		inp.DateStart,
		inp.Discipline,
		inp.Items,
		inp.Status,
		); err != nil {
		c.Err()
		return
	}

	c.Status(http.StatusOK)
}