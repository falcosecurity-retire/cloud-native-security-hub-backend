package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RetrieveAllResourcesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
}

func RegisterOn(router *httprouter.Router) {
	router.GET("/resources", RetrieveAllResourcesHandler)
}
