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

	if http.StatusOK != resp.StatusCode {
		t.Errorf("TestReadMyAPI() StatusCode error: expected %d received %d", http.StatusOK, resp.StatusCode)
	}
}

func TestReadError(t *testing.T) {
	srv := server.Server{
		Router: mux.NewRouter(),
	}
	srv.Routes()

	req, err := http.NewRequest(http.MethodGet, "/error", nil)
	if err != nil {
		t.Errorf("TestReadError() NewRequest(): %s", err.Error())
	}

	w := httptest.NewRecorder()
	srv.Router.ServeHTTP(w, req)

	resp := w.Result()

	if http.StatusInternalServerError != resp.StatusCode {
		t.Errorf("TestReadError() StatusCode error: expected %d received %d", http.StatusInternalServerError, resp.StatusCode)
	}
}

func TestOptions(t *testing.T) {
	srv := server.Server{
		Router: mux.NewRouter(),
	}
	srv.Routes()

	req, err := http.NewRequest(http.MethodOptions, "/myapi", nil)
	if err != nil {
		t.Errorf("TestOptions() NewRequest(): %s", err.Error())
	}

	w := httptest.NewRecorder()
	srv.Router.ServeHTTP(w, req)

	resp := w.Result()

	if resp.Header.Get("Access-Control-Allow-Headers") != "Authorization" {
		t.Errorf("TestOptions(): expected Access-Control-Allow-Headers to be %s received %s", "Authorization", resp.Header.Get("Access-Control-Allow-Headers"))
	}
}
