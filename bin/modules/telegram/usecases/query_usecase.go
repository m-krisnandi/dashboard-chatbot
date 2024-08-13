package usecases

import (
	"context"
	models "dashboard-chatbot/bin/modules/telegram/models/domain"
	"dashboard-chatbot/bin/modules/telegram/repositories/queries"
	"dashboard-chatbot/bin/pkg/utils"
)

type queryUsecase struct {
	postgreQuery queries.QueryPostgre
}

func NewQueryUsecase(postgreQuery queries.QueryPostgre) *queryUsecase {
	return &queryUsecase{
		postgreQuery: postgreQuery,
	}
}

func (q *queryUsecase) GetTourismTypes(ctx context.Context) utils.Result {
	var result utils.Result

	queryPayload := queries.QueryPayload{
		Table: "tourism_types",
		Select: `*`,
		Output: []models.GetTourismTypes{},
	}

	queryRes := <-q.postgreQuery.FindManyBasic(&queryPayload)
	if queryRes.Error != nil {
		queryRes.Data = []models.GetTourismTypes{}
	}

	result.Data = queryRes.Data
	return result
}