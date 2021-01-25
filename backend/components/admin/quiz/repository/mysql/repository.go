package mysql

import (
	"context"
	"database/sql"
	"github.com/abylq/learning-management-system/components/admin/quiz/models"
)

type Quiz struct {
	ID 		 	int64 `json:"id"`
	ParentID 	int64 `json:"parent_id"`
	Title 	 	string `json:"title"`
	CreatedAT  	string
	UpdatedAT  	string
}

type QuizCategoryRepository struct {
	db *sql.DB
}

func NewQuizCategoryRepository(db *sql.DB) *QuizCategoryRepository {
	return &QuizCategoryRepository{
		db: db,
	}
}

func(quc QuizCategoryRepository) CreateCategoryQuiz(ctx context.Context,quiz *models.Quiz) error {
	query := "INSERT quiz SET title=?"
	stmt, err := quc.db.PrepareContext(ctx,query)

	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx,quiz.Title)

	if err != nil {
		return err
	}

	if res != nil {
		return err
	}

	//log.Fatal(lastID)

	return nil
}