package usecase

import (
	"context"
	userModel "github.com/abylq/learning-management-system/components/auth/models"
	"github.com/abylq/learning-management-system/components/common/orders"
	"github.com/abylq/learning-management-system/components/common/orders/models"
)

type OrderUseCase struct {
	orderRepo orders.Repository
}

func NewOrderUseCase(orderRepo orders.Repository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
	}
}

func (o OrderUseCase) CreateOrder(ctx context.Context, user *userModel.User,
	datestart string,discipline int64,
	items,status int,
	) error {
	om := &models.Orders{
		DateStart: datestart,
		Items: items,
		DisciplineID: discipline,
		Status: status,
	}

	return o.orderRepo.CreateOrder(ctx,user,om)
}
