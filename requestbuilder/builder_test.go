package requestbuilder_test

import (
	"fmt"
	"net/http"

	"github.com/anth1y/gurl/requestbuilder"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Builder", func() {
	var (
		requestURL string
		args       []string
		reqBuilder *requestbuilder.RequestBuilder
		request    *http.Request
	)

	Describe("Building a GET request", func() {
		BeforeEach(func() {
			requestURL = "http://someurl.com"
			args = []string{requestURL}
			reqBuilder = requestbuilder.New(args)
		})

		It("returns a http request", func() {
			req, _ := reqBuilder.Build()
			Expect(req).To(BeAssignableToTypeOf(&http.Request{}))
		})

		Describe("basic request", func() {
			BeforeEach(func() {
				request, _ = reqBuilder.Build()
			})

			It("has the provided URL", func() {
				Expect(request.URL.String()).To(Equal(requestURL))
			})

			It("sets the method to GET", func() {
				Expect(request.Method).To(Equal("GET"))
			})
		})

		Context("with an invalid format URL", func() {
			It("returns an error", func() {
				requestURL = "::://someurl.com"
				args = []string{requestURL}
				req, err := requestbuilder.New(args).Build()

				Expect(err.Error()).To(Equal(
					fmt.Sprintf("parse %s: missing protocol scheme", requestURL),
				))
				Expect(req).To(BeNil())
			})
		})
	})

	DescribeTable("request methods",
		func(setMethod, expectedMethod string) {
			requestURL = "http://someurl.com"
			args = []string{"-X", setMethod, requestURL}
			request, _ = requestbuilder.New(args).Build()

			Expect(request.Method).To(Equal(expectedMethod))
			Expect(request.URL.String()).To(Equal(requestURL))
		},
		Entry("GET", "GET", "GET"),
		Entry("POST", "POST", "POST"),
		Entry("PUT", "PUT", "PUT"),
		Entry("DELETE", "DELETE", "DELETE"),
		Entry("HEAD", "HEAD", "HEAD"),
		Entry("PATCH", "PATCH", "PATCH"),
	)
})
