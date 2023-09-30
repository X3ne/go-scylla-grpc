package server

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"scylla-grpc-adapter/config"
	"scylla-grpc-adapter/gen/adapter/v1/adapterv1connect"
	"scylla-grpc-adapter/internal/gateway"
)

func LaunchServer(config *config.Config) {
	api  := http.NewServeMux()
	api.Handle(adapterv1connect.NewAdapterServiceHandler(&gateway.AdaperServer{}))

	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", api))
	log.Println("Server started on port " + config.SERVER.Port)
	err := http.ListenAndServe(
		config.SERVER.Host + ":" + config.SERVER.Port,
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
