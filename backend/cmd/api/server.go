package main

import (
	"github.com/ProstoyVadila/fuckgit/config"
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	config *config.Config
	store  string
}

func NewApp(config *config.Config) (*App, error) {
	app := &App{
		config: config,
		router: gin.New(),
		store:  "db",
	}
	return app, nil
}

func (a *App) Run(addr string) {
	a.router.Run(addr)
}
