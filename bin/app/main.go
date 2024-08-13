package main

import (
	"dashboard-chatbot/bin/config"
	"dashboard-chatbot/bin/pkg/utils"
	"fmt"
	"net/http"

	dashboardHTTP "dashboard-chatbot/bin/modules/telegram/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Validator = utils.NewValidationUtil()
	e.Use(middleware.CORS())

	e.GET("dashboard", func(c echo.Context) error {
		return c.String(http.StatusOK, "This service is running properly")
	})

	dashboardGroup := e.Group("/chatbot")

	dashboardHTTP.New().Mount(dashboardGroup)

	listenerPort := fmt.Sprintf(":%d", config.GlobalEnv.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}