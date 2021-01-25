package mysql

import (
	"context"
	"database/sql"
	userm "github.com/abylq/learning-management-system/components/auth/models"
	"github.com/abylq/learning-management-system/components/common/orders/models"
)

type Order struct {
	ID int64 `json:"id"`
	UserID int `json:"user_id"`
	DateStart string `json:"date_start"`
	Items int `json:"items"`
	DisciplineID string `json:"discipline"`
	Status int `json:"status"`
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r OrderRepository) CreateOrder(ctx context.Context, user *userm.User, om *models.Orders) error {
	query := "INSERT orders SET user_id=?, date_start=?, items=?, discipline=?, status=?"
	stmt, err := r.db.PrepareContext(ctx,query)

	if err != nil {
		return err
	}

	res, err :=stmt.ExecContext(ctx,user.ID,om.DateStart,om.Items,om.DisciplineID,om.Status)

	if err != nil {
		return err
	}
	lastID,err := res.LastInsertId()

	user.ID = string(lastID)

	return nil
}