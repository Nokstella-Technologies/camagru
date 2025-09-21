// Package application for run gin framework
package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func New() *App {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	return &App{
		router: router,
	}
}

func (app *App) Run(port string) error {
	fmt.Printf("Starting server on port %s\n", port)
	return app.router.Run(":" + port)
}
