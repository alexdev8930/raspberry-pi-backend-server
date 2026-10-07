package httptransport

import stdhttp "net/http"

func (h *Handler) Routes() stdhttp.Handler {
	mux := stdhttp.NewServeMux()
	mux.HandleFunc("GET /healthz", h.healthz)
	return mux
}
