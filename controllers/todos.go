package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) error {
	return c.JSON(200, "todos")
}

func AddTodo(c echo.Context) error {
	err, todo := BindTodos(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.String(http.StatusCreated, todo.Description)
}

func UpdateTodo(c echo.Context) error {
	return c.JSON(200, "update todo")
}

func DeleteTodo(c echo.Context) error {
	return c.JSON(200, "delete todo")
}
