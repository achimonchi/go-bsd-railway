package main

import (
	"os"
	"sesi-11/domain/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	api := router.Group("api")

	todo.RegisterRoute(api)

	router.Run(os.Getenv("PORT"))
}
