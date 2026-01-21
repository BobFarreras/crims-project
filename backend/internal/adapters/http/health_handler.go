package http

import (
	"net/http"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
	"github.com/digitaistudios/crims-backend/internal/ports"
)

type healthResponse struct {
	Status     string `json:"status"`
	PocketBase string `json:"pocketbase"`
}

func NewHealthHandler(pbClient ports.PocketBaseClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := pbClient.Ping(r.Context()); err != nil {
			web.RespondError(w, http.StatusServiceUnavailable, "pocketbase unavailable", "pocketbase_unavailable")
			return
		}

		web.RespondJSON(w, http.StatusOK, healthResponse{
			Status:     "healthy",
			PocketBase: "ok",
		})
	}
}
