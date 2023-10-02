package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"scylla-grpc-adapter/config"
	"scylla-grpc-adapter/gen/adapter/v1/adapterv1connect"
	"scylla-grpc-adapter/gen/users/v1/usersv1connect"
	"scylla-grpc-adapter/internal/gateway"
)

type Server struct {}

func LaunchServer(cfg *config.Config) {
	api  := http.NewServeMux()
	api.Handle(adapterv1connect.NewAdapterServiceHandler(&gateway.AdaperServer{}))
	api.Handle(usersv1connect.NewUsersServiceHandler(&gateway.UsersServer{}))

	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", api))

	srv := &http.Server{
		Addr: cfg.SERVER.Host + ":" + cfg.SERVER.Port,
		Handler: h2c.NewHandler(
			mux,
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024,
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	log.Printf("Starting HTTP server on %s", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP listen and serve: %v", err)
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown: %v", err)
	}
}
