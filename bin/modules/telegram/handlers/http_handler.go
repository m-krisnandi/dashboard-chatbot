package handlers

import (
	"context"
	"dashboard-chatbot/bin/modules/telegram/repositories/queries"
	"dashboard-chatbot/bin/modules/telegram/usecases"
	"dashboard-chatbot/bin/pkg/database"
	"dashboard-chatbot/bin/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTPHandler struct
type HTTPHandler struct {
	queryUsecase usecases.QueryUsecase
}

// New initiation
func New() *HTTPHandler {
	postgreDb := database.InitPostgre(context.Background())

	postgreQuery := queries.NewPostgreQuery(postgreDb)
	queryUsecase := usecases.NewQueryUsecase(postgreQuery)

	return &HTTPHandler{
		queryUsecase: queryUsecase,
	}
}

// Mount function
func (u *HTTPHandler) Mount(g *echo.Group) {
	g.GET("/v1/dashboard/tourism-types", u.getTourismTypes)
}

func (u *HTTPHandler) getTourismTypes(c echo.Context) error {
	result := u.queryUsecase.GetTourismTypes(c.Request().Context())
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Tourism: Get List of Tourism Types", http.StatusOK, c)
}
