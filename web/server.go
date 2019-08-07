package web

import (
	"cloud-native-visibility-hub/web/api"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	api.RegisterOn(router)
	return router
}
