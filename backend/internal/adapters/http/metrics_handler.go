package http

import (
	"net/http"
	"time"

	"github.com/digitaistudios/crims-backend/internal/platform/web"
)

var startTime = time.Now()

type metricsResponse struct {
	UptimeSeconds int64 `json:"uptime_seconds"`
}

func NewMetricsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uptime := time.Since(startTime).Seconds()
		web.RespondJSON(w, http.StatusOK, metricsResponse{UptimeSeconds: int64(uptime)})
	}
}
