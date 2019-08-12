package web

import (
	"cloud-native-visibility-hub/pkg/resource"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRetrieveAllResourcesHandlerReturnsHTTPOk(t *testing.T) {
	request, _ := http.NewRequest("GET", "/resources", nil)
	recorder := httptest.NewRecorder()
	os.Setenv("RESOURCES_PATH", "../test/fixtures")
	router := NewRouter()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestRetrieveAllResourcesHandlerReturnsResourcesSerializedAsJSON(t *testing.T) {
	request, _ := http.NewRequest("GET", "/resources", nil)
	recorder := httptest.NewRecorder()
	path := "../test/fixtures"
	repo, err := resource.NewFile(path)
	assert.Equal(t, nil, err)
	resources, _ := repo.All()

	router := NewRouter()
	router.ServeHTTP(recorder, request)

	var result []resource.Resource
	body, _ := ioutil.ReadAll(recorder.Body)
	json.Unmarshal([]byte(body), &result)

	assert.Equal(t, resources, result)
}

func TestRetrieveAllResourcesHandlerReturnsAJSONResponse(t *testing.T) {
	request, _ := http.NewRequest("GET", "/resources", nil)
	recorder := httptest.NewRecorder()
	os.Setenv("RESOURCES_PATH", "../test/fixtures")
	router := NewRouter()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, jsonHeader(), recorder.HeaderMap)
}

func jsonHeader() http.Header {
	jsonHeader := http.Header{}
	jsonHeader.Set("Content-Type", "application/json")
	return jsonHeader
}
