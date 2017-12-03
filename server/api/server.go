package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	API_PREFIX = "/api"
)

func Serve(addr, endpoint string) error {
	mux := http.NewServeMux()

	gmux, err := newGateway(endpoint)
	if err != nil {
		return err
	}
	mux.Handle("/rpc", gmux)

	mux.Handle("/api", apiServer())

	mux.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(addr, mux)
}

func apiServer() http.Handler {
	g := gin.New()
	g.Use(Logger(), Recovery())
	return g
}
