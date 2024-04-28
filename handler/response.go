package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendSuccess(ctx *gin.Context, op string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"response": fmt.Sprint(op),
	})
}
