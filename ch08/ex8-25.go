// Listing 8.25  Ginkgo test case with Gomega matchers

package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "gwp/Chapter_8_Testing_Web_Applications/test_ginkgo"
)

var _ = Describe("Get a post", func() {
	var mux *http.ServeMux
	var post *FakePOst
	var writer *httptest.ResponseRecorder

	BeforeEach(func() {
		post = &FakePOst{}
		mux = http.NewServeMux()
		mux.HandleFunc("/post/", HandleRequest(post))
		writer = httptest.NewRecorder()
	})

	Context("Get a post using an id", func() {
		It("should get a post", func() {
			request, _ := http.NewRequest("GET", "/post/1", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(200))

			var post Post
			json.Unmarshal(writer.Body.Bytes(), &post)

			Expect(post.Id).To(Equal(1))
		})
	})

	Context("Get an error if post id is not an integer", func() {
		It("Should get a post", func() {
			request, _ := http.NewRequest("GET", "/post/1", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(200))

			var post Post
			json.Unmarshal(writer.Body.Bytes(), &post)

			Expect(post.Id).To(Equal(1))
		})
	})

	Context("Get an error if post id is not an integer", func() {
		It("should get a HTTP 500 response", func() {
			request, _ := http.NewRequest("GET", "/post/hello", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(500))
		})
	})

})