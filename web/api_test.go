package web

import (
	"cloud-native-visibility-hub/pkg/resource"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRetrieveAllResourcesHandlerReturnsHTTPOk(t *testing.T) {
	request, _ := http.NewRequest("GET", "/resources", nil)
	recorder := httptest.NewRecorder()

	router := NewRouter()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestRetrieveAllResourcesHandlerReturnsResourcesSerializedAsJSON(t *testing.T) {
	request, _ := http.NewRequest("GET", "/resources", nil)
	recorder := httptest.NewRecorder()

	router := NewRouter()
	router.ServeHTTP(recorder, request)

	var result []resource.Resource
	body, _ := ioutil.ReadAll(recorder.Body)
	json.Unmarshal([]byte(body), &result)

	assert.Equal(t, []resource.Resource{
		{
			Name: "Falco profile for Nginx",
		},
		{
			Name: "Grafana Dashboard for Traefik",
		},
	}, result)
}

func TestRetrieveAllResourcesHandlerReturnsAJSONResponse(t *testing.T) {
	request, _ := http.NewRequest("GET", "/resources", nil)
	recorder := httptest.NewRecorder()

	router := NewRouter()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, jsonHeader(), recorder.HeaderMap)
}

func jsonHeader() http.Header {
	jsonHeader := http.Header{}
	jsonHeader.Set("Content-Type", "application/json")
	return jsonHeader
}
