package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func newHandler() handler {
	return handler{}
}

func (h handler) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
