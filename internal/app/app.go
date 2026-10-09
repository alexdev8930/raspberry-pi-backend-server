package app

import (
    "log"
    "net/http"

    "rpi-backend.local/internal/config"
    "rpi-backend.local/internal/service"
    httptransport "rpi-backend.local/internal/transport/http"
)

type App struct {
	cfg *config.Config
	handler http.Handler
}

func New(cfg *config.Config) *App {
	svc := service.New()
	h := httptransport.NewHandler(svc)

	return &App{
		cfg: cfg,
		handler: h.Routes(),
	}
}

func (a *App) Run() error {
	server := &http.Server{
        Addr:    a.cfg.Port,
        Handler: a.handler,
    }

    log.Println("server listening on", a.cfg.Port)
    return server.ListenAndServe()
}


