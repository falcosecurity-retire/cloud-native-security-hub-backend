package web_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
)

var _ = Describe("HTTP HealthCheck", func() {
	Context("GET /resources", func() {
		It("returns OK", func() {
			response := doGetRequest("/health")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})
	})
})
