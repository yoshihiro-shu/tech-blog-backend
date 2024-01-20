package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metricsRouter struct{}

func (m *metricsRouter) SetRouters(router *mux.Router) {

	router.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)
}
