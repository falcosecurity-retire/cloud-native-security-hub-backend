package web

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func NewRouter() http.Handler {
	router := httprouter.New()
	registerOn(router, nil)

	return cors.Default().Handler(router)
}

func NewRouterWithLogger(logger *log.Logger) http.Handler {
	router := httprouter.New()
	registerOn(router, logger)

	return cors.Default().Handler(router)
}

func registerOn(router *httprouter.Router, logger *log.Logger) {
	h := NewHandlerRepository(logger)
	router.GET("/resources", h.retrieveAllResourcesHandler)

	router.GET("/resources/:kind/:resource", h.retrieveOneResourcesHandler)
	router.GET("/resources/:kind/:resource/version/:version", h.retrieveOneResourceByVersionHandler)

	router.GET("/resources/:kind/:resource/custom-rules.yaml", h.retrieveFalcoRulesForHelmChartHandler)
	router.GET("/resources/:kind/:resource/version/:version/custom-rules.yaml", h.retrieveFalcoRulesForHelmChartByVersionHandler)

	router.GET("/vendors", h.retrieveAllVendorsHandler)
	router.GET("/vendors/:vendor", h.retrieveOneVendorsHandler)
	router.GET("/vendors/:vendor/resources", h.retrieveAllResourcesFromVendorHandler)

	router.GET("/health", h.healthCheckHandler)
	router.NotFound = h.notFound()
}
