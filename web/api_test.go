package web

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRetrieveAllResourcesHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/resources", nil)
	recorder := httptest.NewRecorder()

	router := NewRouter()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, recorder.Code, http.StatusOK)
}
