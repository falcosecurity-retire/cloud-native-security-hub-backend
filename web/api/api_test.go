package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRetrieveAllResourcesHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/resources", nil)
	recorder := httptest.NewRecorder()

	router := newRouter()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, recorder.Code, http.StatusOK)
}

func newRouter() *httprouter.Router {
	router := httprouter.New()
	RegisterOn(router)
	return router
}
