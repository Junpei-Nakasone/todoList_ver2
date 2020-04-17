package api001

import (
	"net/http"

	"github.com/labstack/echo"
)

func ShowTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, getTodos())
}

type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name`
}

type TodoCollection struct {
	Todos []Todo `json:"items"`
}

func getTodos() (tc TodoCollection) {
	tc = TodoCollection{
		[]Todo{
			{1, "cleaning"},
			{2, "cooking"},
		},
	}
	return
}
