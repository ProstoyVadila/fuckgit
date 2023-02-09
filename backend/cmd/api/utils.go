package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// finishSuccessfully is a helper function for finishing request with success.
func finishSuccessfully(ctx *gin.Context, payload any) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"payload": payload,
	})
}

// finishWithError is a helper function for finishing request with error and http.StatusInternalServerError as a default status.
func finishWithError(ctx *gin.Context, err error, httpStatuses ...int) {
	var httpStatus int
	if len(httpStatuses) > 0 {
		httpStatus = httpStatuses[0]
	} else {
		httpStatus = http.StatusInternalServerError
	}
	ctx.JSON(httpStatus, gin.H{
		"status": "error",
		"error":  err.Error(),
	})
}
