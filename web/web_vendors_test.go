package web_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
)

var _ = Describe("HTTP API for vendors", func() {
	Context("GET /vendors", func() {
		It("returns OK", func() {
			response := doGetRequest("/vendors")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/vendors")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})
	})

	Context("GET /vendors/:name", func() {
		It("returns OK", func() {
			response := doGetRequest("/vendors/apache")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/vendors/apache")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		PContext("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/vendors/non-existent")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	Context("GET /vendors/:name/resources", func() {
		It("returns OK", func() {
			response := doGetRequest("/vendors/apache/resources")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/vendors/apache/resources")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		PContext("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/vendors/non-existent/resources")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})
})
