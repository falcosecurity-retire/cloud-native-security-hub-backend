package web

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join("testdata", "fixture.sql") // relative path
	data, err := ioutil.ReadFile(path)

	_, err = db.Exec(`TRUNCATE TABLE latest_security_resources;
					  TRUNCATE TABLE security_resources;
					  TRUNCATE TABLE vendors;
					  TRUNCATE TABLE schema_migrations`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(data))
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestRetrieveAllResourcesHandlerReturnsHTTPOk(t *testing.T) {
	testRetrieveAllReturnsHTTPOk(t, "/resources")
}

func TestRetrieveOneResourceHandlerReturnsHTTPOk(t *testing.T) {
	apacheID := "apache"
	testRetrieveAllReturnsHTTPOk(t, "/resources/"+apacheID)
}

func TestRetrieveFalcoRulesForHelmChartHandlerReturnsHTTPOk(t *testing.T) {
	apacheID := "apache"
	testRetrieveAllReturnsHTTPOk(t, "/resources/"+apacheID+"/custom-rules.yaml")
}

func TestRetrieveAllVendorsHandlerReturnsHTTPOk(t *testing.T) {
	testRetrieveAllReturnsHTTPOk(t, "/vendors")
}

func TestRetrieveOneVendorsHandlerReturnsHTTPOk(t *testing.T) {
	testRetrieveAllReturnsHTTPOk(t, "/vendors/apache")
}

func TestRetrieveAllResourcesFromVendorHandlerReturnsHTTPOk(t *testing.T) {
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
	apacheID := "apache"
	request, _ := http.NewRequest("GET", "/resources/"+apacheID+"/custom-rules.yaml", nil)
	recorder := httptest.NewRecorder()

	router := NewRouter()
	router.ServeHTTP(recorder, request)
	assert.Equal(t, "application/x-yaml", recorder.HeaderMap["Content-Type"][0])
}

func TestLoggerIsLogging(t *testing.T) {
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
