package util

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/steinfletcher/apitest"
)

// BasicCheck 基本的な項目をテストする
func BasicCheck(t *testing.T, apiURL string, app *echo.Echo, allowedMethod RequestMethod) {
	t.Helper()
	if allowedMethod&Get != Get {
		t.Run("Getメソッドは許可されていません", func(t *testing.T) {
			apitest.New().
				Handler(app).
				Get(apiURL).
				Expect(t).
				Status(http.StatusMethodNotAllowed).
				End()
		})
	}
	if allowedMethod&Post != Post {
		t.Run("Postメソッドは許可されていません", func(t *testing.T) {
			apitest.New().
				Handler(app).
				Post(apiURL).
				Expect(t).
				Status(http.StatusMethodNotAllowed).
				End()
		})
	}
	if allowedMethod&Put != Put {
		t.Run("Putメソッドは許可されていません", func(t *testing.T) {
			apitest.New().
				Handler(app).
				Post(apiURL).
				Expect(t).
				Status(http.StatusMethodNotAllowed).
				End()
		})
	}
	if allowedMethod&Delete != Delete {
		t.Run("Deleteメソッドは許可されていません", func(t *testing.T) {
			apitest.New().
				Handler(app).
				Post(apiURL).
				Expect(t).
				Status(http.StatusMethodNotAllowed).
				End()
		})
	}
}
