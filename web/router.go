package web

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	registerOn(router)
	return router
}

func registerOn(router *httprouter.Router) {
	h := NewHandlerRepository()
	router.GET("/resources", h.retrieveAllResourcesHandler)
	router.GET("/resources/:hash", h.retrieveOneResourcesHandler)
	router.GET("/vendors", h.retrieveAllVendorsHandler)
	router.GET("/vendors/:vendor", h.retrieveOneVendorsHandler)
	router.GET("/vendors/:vendor/resources", h.retrieveAllResourcesFromVendorHandler)
}
