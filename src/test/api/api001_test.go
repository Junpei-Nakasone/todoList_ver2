package api

import (
	"testing"

	"github.com/todoList_ver2/environment/db"
	"github.com/todoList_ver2/environment/router"
	"github.com/todoList_ver2/test/util"
)

func TestHandler_api_001(t *testing.T) {
	db := db.CreateDBConnection()
	defer db.Close()

	e := router.NewRouter()
	util.BasicCheck(t, "/todos", e, util.Get)
}
