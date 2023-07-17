package todo

import "github.com/gin-gonic/gin"

func RegisterRoute(router *gin.RouterGroup) {
	h := newHandler()

	router.GET("/todos/health", h.healthCheck)
}
