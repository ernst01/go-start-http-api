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

		if resp.ID > (rmax / 2) {
			response.SendJSONError(w, http.StatusInternalServerError, "http://example.com/helper_url", "Something went wrong: %d is greater than %d", resp.ID, rmax/2)
		} else {
			response.SendJSONSuccess(w, http.StatusOK, resp)
		}
	}
}
