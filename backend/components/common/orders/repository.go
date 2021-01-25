package orders

import (
	"context"
	auth "github.com/abylq/learning-management-system/components/auth/models"
	"github.com/abylq/learning-management-system/components/common/orders/models"
)

type Repository interface {
	CreateOrder(ctx context.Context, user *auth.User, order *models.Orders) error

}
