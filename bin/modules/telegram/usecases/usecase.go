package usecases

import (
	"context"
	"dashboard-chatbot/bin/pkg/utils"
)

type QueryUsecase interface {
	GetTourismTypes(ctx context.Context) utils.Result
}