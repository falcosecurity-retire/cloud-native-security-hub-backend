package web_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
)

var _ = Describe("HTTP API for resources", func() {
	Context("GET /resources", func() {
		It("returns OK", func() {
			response := doGetRequest("/resources")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/resources")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})
	})

	Context("GET /resources/:name", func() {
		It("returns OK", func() {
			response := doGetRequest("/resources/apache")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/resources/apache")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		PContext("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/non-existent")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	Context("GET /resources/:name/custom-rules.yaml", func() {
		It("returns OK", func() {
			response := doGetRequest("/resources/apache/custom-rules.yaml")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an YAML response", func() {
			response := doGetRequest("/resources/apache/custom-rules.yaml")

			Expect(response.Header.Get("Content-Type"), "application/x-yaml")
		})

		PContext("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/non-existent/custom-rules.yaml")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	Context("GET /resources/:name/version/:version", func() {
		It("returns OK", func() {
			response := doGetRequest("/resources/apache/version/1.0.0")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/resources/apache")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		PContext("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/non-existent/version/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})
})
