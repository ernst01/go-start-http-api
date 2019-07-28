package main

import (
	"log"
	"net/http"

	"github.com/ernst01/go-start-http-api/internal/myapi"
	"github.com/gorilla/mux"
)

func main() {
	srv := myapi.Server{
		Router: mux.NewRouter(),
	}

	srv.Routes()

	http.Handle("/", srv.Router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
