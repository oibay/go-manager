package quiz

import (
	"context"
	"github.com/abylq/learning-management-system/components/admin/quiz/models"
)

type Repository interface {
	CreateCategoryQuiz(ctx context.Context, quiz *models.Quiz) error

}