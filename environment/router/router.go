package router

import (
	"github.com/labstack/echo"
	api001 "github.com/todoList_ver2/api/api001/handler"
	api002 "github.com/todoList_ver2/api/api002/handler"
)

// NewRouter returns router
func NewRouter() *echo.Echo {
	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/todos", api001.ShowTodos)
	e.POST("/addTodo", api002.AddTodo)

	return e
}
