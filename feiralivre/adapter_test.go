package feiralivre

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

var (
	mockDB = map[string]interface{}{
		"name": "Bruno",
	}
)

func TestWelcomeHTTPHandler(t *testing.T) {
	e := echo.New()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	c := e.NewContext(req, rec)

	c.SetParamNames("name")
	c.SetParamValues("Bruno")

	resp := rec.Result()

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned invalid status code: got %v want %v", status, http.StatusOK)
	}
}
