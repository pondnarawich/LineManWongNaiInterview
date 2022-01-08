package app

import (
	"LineManWongNaiInternShip/miniproject/app/user"
	"github.com/gin-gonic/gin"

	userService "LineManWongNaiInternShip/miniproject/service/user"
)

type App struct {
	user *user.Controller
}

func New(userService userService.Service) *App {
	return &App{
		user: user.New(userService),
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	adminRoute := router.Group("/covid")
	{
		adminRoute.GET("/summary", app.user.Summary)
	}

	return app
}
