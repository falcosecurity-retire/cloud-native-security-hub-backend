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
	testRetrieveAllReturnsHTTPOk(t, "/resources", "../test/fixtures/resources")
}

func TestRetrieveOneResourceHandlerReturnsHTTPOk(t *testing.T) {
	apacheHash := "ijdncmo23ix4j5cgdwwh"
	testRetrieveAllReturnsHTTPOk(t, "/resources/"+apacheHash, "../test/fixtures/resources")
}

func TestRetrieveAllVendorsHandlerReturnsHTTPOk(t *testing.T) {
	testRetrieveAllReturnsHTTPOk(t, "/vendors", "../test/fixtures/vendors")
}

func TestRetrieveOneVendorsHandlerReturnsHTTPOk(t *testing.T) {
	testRetrieveAllReturnsHTTPOk(t, "/vendors/apache", "../test/fixtures/vendors")
}

func testRetrieveAllReturnsHTTPOk(t *testing.T, path, fixturesPath string) {
	request, _ := http.NewRequest("GET", path, nil)
	recorder := httptest.NewRecorder()
	os.Setenv("RESOURCES_PATH", "../test/fixtures/resources")
	os.Setenv("VENDOR_PATH", "../test/fixtures/vendors")
	router := NewRouter()
	router.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestRetrieveAllResourcesHandlerReturnsResourcesSerializedAsJSON(t *testing.T) {
	testRetrieveallSerializedAsJSON(t, "/resources", "../test/fixtures/resources")
}

func TestRetrieveAllVendorsHandlerReturnsResourcesSerializedAsJSON(t *testing.T) {
	testRetrieveallSerializedAsJSON(t, "/vendors", "../test/fixtures/vendors")
}

func testRetrieveallSerializedAsJSON(t *testing.T, urlPath, fixturesPath string) {
	request, _ := http.NewRequest("GET", urlPath, nil)
	recorder := httptest.NewRecorder()
	os.Setenv("RESOURCES_PATH", "../test/fixtures/resources")
	os.Setenv("VENDOR_PATH", "../test/fixtures/vendors")
	repo, err := resource.NewFile(fixturesPath)
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
	testRetrieveAllHandlerReturnsAJSONResponse(t, "/resources")
}
func TestRetrieveAllVendorHandlerReturnsAJSONResponse(t *testing.T) {
	testRetrieveAllHandlerReturnsAJSONResponse(t, "/vendors")
}

func testRetrieveAllHandlerReturnsAJSONResponse(t *testing.T, urlPath string) {
	request, _ := http.NewRequest("GET", urlPath, nil)
	recorder := httptest.NewRecorder()
	os.Setenv("RESOURCES_PATH", "../test/fixtures/resources")
	os.Setenv("VENDOR_PATH", "../test/fixtures/vendors")
	router := NewRouter()
	router.ServeHTTP(recorder, request)
	assert.Equal(t, jsonHeader(), recorder.HeaderMap)
}

func jsonHeader() http.Header {
	jsonHeader := http.Header{}
	jsonHeader.Set("Content-Type", "application/json")
	return jsonHeader
}
