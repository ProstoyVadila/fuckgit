package main

import (
	"github.com/ProstoyVadila/fuckgit/config"
	"github.com/ProstoyVadila/fuckgit/repository"
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	config *config.Config
	store  repository.Store
}

func NewApp(config *config.Config) (*App, error) {
	app := &App{
		config: config,
		store:  repository.New(),
		router: gin.New(),
	}

	app.setMiddlewares()
	app.setRoutes()

	return app, nil
}

func (a *App) Run() {
	a.router.Run(a.config.Addr)
}
