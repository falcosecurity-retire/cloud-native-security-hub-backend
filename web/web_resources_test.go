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

	Context("GET /resources/:kind/:name", func() {
		It("returns OK", func() {
			response := doGetRequest("/resources/falco-rules/apache")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/resources/falco-rules/apache")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		PContext("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/falco-rules/non-existent")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when kind does not exist", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/foo/apache")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	Context("GET /resources/:kind/:name/custom-rules.yaml", func() {
		It("returns OK", func() {
			response := doGetRequest("/resources/falco-rules/apache/custom-rules.yaml")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an YAML response", func() {
			response := doGetRequest("/resources/falco-rules/apache/custom-rules.yaml")

			Expect(response.Header.Get("Content-Type"), "application/x-yaml")
		})

		PContext("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/non-existent/custom-rules.yaml")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when is not falco-rules", func() {
			It("returns a BADREQUEST", func() {
				response := doGetRequest("/resources/foo/apache/custom-rules.yaml")

				Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
	})

	Context("GET /resources/:kind/:name/version/:version", func() {
		It("returns OK", func() {
			response := doGetRequest("/resources/falco-rules/apache/version/1.0.0")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/resources/falco-rules/apache")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		PContext("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/falco-rules/non-existent/version/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when kind does not exist", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/foo/apache/version/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	Context("GET /resources/:kind/:name/version/:version/custom-rules.yaml", func() {
		It("returns OK", func() {
			response := doGetRequest("/resources/falco-rules/apache/version/1.0.0/custom-rules.yaml")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an YAML response", func() {
			response := doGetRequest("/resources/falco-rules/apache/version/1.0.0/custom-rules.yaml")

			Expect(response.Header.Get("Content-Type"), "application/x-yaml")
		})

		PContext("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/falco-rules/apache/version/4.0.0/custom-rules.yaml")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when is not falco-rules", func() {
			It("returns a BADREQUEST", func() {
				response := doGetRequest("/resources/foo/apache/version/1.0.0/custom-rules.yaml")

				Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
	})

})
