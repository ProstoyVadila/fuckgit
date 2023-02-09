package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) ping(ctx *gin.Context) {
	finishSuccessfully(ctx, "pong")
}

func (a *App) questions(ctx *gin.Context) {
	questions, err := a.store.Questions()
	if err != nil {
		finishWithError(ctx, err, http.StatusNotFound)
		return
	}
	finishSuccessfully(ctx, questions)
}
