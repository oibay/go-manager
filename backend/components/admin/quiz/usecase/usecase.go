package usecase

import (
	"context"
	"github.com/abylq/learning-management-system/components/admin/quiz"
	"github.com/abylq/learning-management-system/components/admin/quiz/models"

)

type QuizCategoryUseCase struct {
	qc quiz.Repository
}

func NewOrderUseCase(qc quiz.Repository) *QuizCategoryUseCase {
	return &QuizCategoryUseCase{
		qc: qc,
	}
}
// quc :QuizCategory
func (quc QuizCategoryUseCase) CreateCategoryQuiz(ctx context.Context,title string) error {
	qcategory := &models.Quiz{
		Title: title,
	}

	return quc.qc.CreateCategoryQuiz(ctx,qcategory)
}