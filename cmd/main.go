package main

import (
	"github.com/devcael/go-todos/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", controllers.GetTodos)
	e.POST("/", controllers.AddTodo)
	e.PUT("/todos/:id", controllers.UpdateTodo)
	e.DELETE("/todos/:id", controllers.DeleteTodo)
	// Comentario pra teste do actions

	e.Logger.Fatal(e.Start(":4040"))
}
