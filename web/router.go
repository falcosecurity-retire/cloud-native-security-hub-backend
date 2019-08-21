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
	router.GET("/resources", retrieveAllResourcesHandler)
	router.GET("/resources/:hash", retrieveOneResourcesHandler)
	router.GET("/vendors", retrieveAllVendorsHandler)
	router.GET("/vendors/:vendor", retrieveOneVendorsHandler)
}
