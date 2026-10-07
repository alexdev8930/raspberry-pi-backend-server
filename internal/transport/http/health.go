package httptransport

import (
	"encoding/json"
	stdhttp "net/http"
)

func (h *Handler) healthz(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(stdhttp.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
