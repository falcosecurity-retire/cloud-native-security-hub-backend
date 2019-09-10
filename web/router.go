package web

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
)

func NewRouter() http.Handler {
	router := httprouter.New()
	registerOn(router)

	return cors.Default().Handler(router)
}

func registerOn(router *httprouter.Router) {
	h := NewHandlerRepository()
	router.GET("/resources", h.retrieveAllResourcesHandler)
	router.GET("/resources/:hash", h.retrieveOneResourcesHandler)
	router.GET("/resources/:hash/raw.yaml", h.retrieveOneResourcesRawHandler)
	router.GET("/vendors", h.retrieveAllVendorsHandler)
	router.GET("/vendors/:vendor", h.retrieveOneVendorsHandler)
	router.GET("/vendors/:vendor/resources", h.retrieveAllResourcesFromVendorHandler)
}
