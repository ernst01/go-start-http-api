package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernst01/go-start-http-api/internal/server"
	"github.com/gorilla/mux"
)

func TestReadMyAPI(t *testing.T) {
	srv := server.Server{
		Router: mux.NewRouter(),
	}
	srv.Routes()

	req, err := http.NewRequest(http.MethodGet, "/myapi", nil)
	if err != nil {
		t.Errorf("TestReadMyAPI() NewRequest(): %s", err.Error())
	}

	w := httptest.NewRecorder()
	srv.Router.ServeHTTP(w, req)

	resp := w.Result()

	if http.StatusOK != resp.StatusCode && http.StatusInternalServerError != resp.StatusCode {
		t.Errorf("TestReadMyAPI() StatusCode error: expected %d or %d received %d", http.StatusOK, http.StatusInternalServerError, resp.StatusCode)
	}
}
