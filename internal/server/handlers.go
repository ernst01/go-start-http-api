package server

import (
	"math/rand"
	"net/http"

	"github.com/ernst01/common/pkg/response"
)

const rmax = 100000

func (s *Server) handleReadMyAPI() http.HandlerFunc {
	type MyAPIResponse struct {
		ID      int    `json:"id"`
		Message string `json:"message"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &MyAPIResponse{
			ID:      rand.Intn(rmax),
			Message: "Well done Sparky!",
		}

		response.SendJSONSuccess(w, http.StatusOK, resp)
	}
}

func (s *Server) handleReadError() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.SendJSONError(w, http.StatusInternalServerError, "http://example.com/helper_url", "Something went wrong")
	}
}
