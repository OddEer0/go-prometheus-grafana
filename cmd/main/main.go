package main

import (
	"grafana-dashboard/internal/cache"
	"grafana-dashboard/internal/router"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))
	c := cache.New[struct{}]()
	r := router.New(c)
	monR := router.NewMon()

	serveCh := make(chan struct{})
	monCh := make(chan struct{})

	go func() {
		_ = http.ListenAndServe(":4000", r)
		serveCh <- struct{}{}
	}()

	go func() {
		_ = http.ListenAndServe(":4010", monR)
		monCh <- struct{}{}
	}()

	<-serveCh
	<-monCh
}
