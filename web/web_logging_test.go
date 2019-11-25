package web_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/falcosecurity/cloud-native-security-hub/web"
)

var _ = Describe("HTTP Server Logging", func() {
	It("is logging requests", func() {
		request, _ := http.NewRequest("GET", "/resources/apache/custom-rules.yaml", nil)
		recorder := httptest.NewRecorder()

		buff := &bytes.Buffer{}
		router := web.NewRouterWithLogger(log.New(buff, "", 0))
		router.ServeHTTP(recorder, request)

		Expect("200 [] GET /resources/apache/custom-rules.yaml\n").To(Equal(buff.String()))
	})
})
