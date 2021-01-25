package orders

import (
	"context"
	"github.com/abylq/learning-management-system/components/auth/models"
)

type UseCase interface {
	CreateOrder(ctx context.Context, user *models.User,
		datestart string,discipline int64,
		items,status int,) error
}
