package router

import (
	"github.com/labstack/echo"
	"github.com/todoList_ver2/api/api001"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/todos", api001.ShowTodos)

	return e
}
