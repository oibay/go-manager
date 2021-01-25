package quiz

import (
	"context"
)

type UseCase interface {
	CreateCategoryQuiz(ctx context.Context, title string) error
}