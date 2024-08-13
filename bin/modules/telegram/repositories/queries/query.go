package queries

import "dashboard-chatbot/bin/pkg/utils"

type QueryPostgre interface {
	FindOne(payload *QueryPayload) <-chan utils.Result
	FindMany(payload *QueryPayload) <-chan utils.Result
	FindManyBasic(payload *QueryPayload) <-chan utils.Result
}