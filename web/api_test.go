package web

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("RESOURCES_PATH", "../test/fixtures/resources")
	os.Setenv("VENDOR_PATH", "../test/fixtures/vendors")

	m.Run()
}

func TestRetrieveAllResourcesHandlerReturnsHTTPOk(t *testing.T) {
	testRetrieveAllReturnsHTTPOk(t, "/resources")
}

func TestRetrieveOneResourceHandlerReturnsHTTPOk(t *testing.T) {
	t.Skip("Disabled until fixture handling")
	apacheID := "apache"
	testRetrieveAllReturnsHTTPOk(t, "/resources/"+apacheID)
}

func TestRetrieveFalcoRulesForHelmChartHandlerReturnsHTTPOk(t *testing.T) {
	t.Skip("Disabled until fixture handling")
	apacheID := "apache"
	testRetrieveAllReturnsHTTPOk(t, "/resources/"+apacheID+"/custom-rules.yaml")
}

func TestRetrieveAllVendorsHandlerReturnsHTTPOk(t *testing.T) {
	testRetrieveAllReturnsHTTPOk(t, "/vendors")
}

func TestRetrieveOneVendorsHandlerReturnsHTTPOk(t *testing.T) {
	t.Skip("Disabled until fixture handling")
	testRetrieveAllReturnsHTTPOk(t, "/vendors/apache")
}

func TestRetrieveAllResourcesFromVendorHandlerReturnsHTTPOk(t *testing.T) {
	t.Skip("Disabled until fixture handling")
	testRetrieveAllReturnsHTTPOk(t, "/vendors/apache/resources")
}

func testRetrieveAllReturnsHTTPOk(t *testing.T, path string) {
	request, _ := http.NewRequest("GET", path, nil)
	recorder := httptest.NewRecorder()

	router := NewRouter()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
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

	router := NewRouter()
	router.ServeHTTP(recorder, request)
	assert.Equal(t, "application/json", recorder.HeaderMap["Content-Type"][0])
}

func TestRetrieveFalcoRulesForHelmChartReturnsAYAMLResponse(t *testing.T) {
	t.Skip("Disabled until fixture handling")

	apacheID := "apache"
	request, _ := http.NewRequest("GET", "/resources/"+apacheID+"/custom-rules.yaml", nil)
	recorder := httptest.NewRecorder()

	router := NewRouter()
	router.ServeHTTP(recorder, request)
	assert.Equal(t, "application/x-yaml", recorder.HeaderMap["Content-Type"][0])
}

func TestLoggerIsLogging(t *testing.T) {
	t.Skip("Disabled until fixture handling")
	apacheID := "apache"
	url := "/resources/" + apacheID + "/custom-rules.yaml"
	request, _ := http.NewRequest("GET", url, nil)
	recorder := httptest.NewRecorder()

	buff := &bytes.Buffer{}
	router := NewRouterWithLogger(log.New(buff, "", 0))
	router.ServeHTTP(recorder, request)

	expectedLog := fmt.Sprintf("200 [] GET %s\n", url)
	assert.Equal(t, expectedLog, buff.String())
}

func TestHealthCheckEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/health", nil)
	recorder := httptest.NewRecorder()

	router := NewRouter()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
