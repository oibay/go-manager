package mysql

import (
	"context"
	"database/sql"
	"github.com/abylq/learning-management-system/components/auth"
	"github.com/abylq/learning-management-system/components/auth/models"
)

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT users SET username=?, password=?"
	stmt, err := r.db.PrepareContext(ctx,query)

	if err != nil {
		return err
	}

	res, err :=stmt.ExecContext(ctx,user.Username,user.Password)

	if err != nil {
		return err
	}
	lastID,err := res.LastInsertId()

	user.ID = string(lastID)

	return err
}

func (r UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	query := "SELECT id,username,password FROM `users` WHERE username=? and password=?"
	rows, err := r.fetch(ctx,query,username,password)

	if err != nil {
		return nil, err
	}
	payload := &models.User{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, auth.ErrUserNotFound
	}

	return payload, nil
}

func (r *UserRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.User, 0)
	for rows.Next() {
		data := new(models.User)

		err := rows.Scan(
			&data.ID,
			&data.Username,
			&data.Password,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

