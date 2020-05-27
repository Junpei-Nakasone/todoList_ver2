package api

import (
	"testing"

	"github.com/todoList_ver2/environment/db"
	"github.com/todoList_ver2/environment/router"
)

func TestHandler_api_001(t *testing.T) {
	db := db.CreateDBConnection()
	defer db.Close()

	e := router.NewRouter()
	BasicCheck(t, "/todos", e, Get)
}
