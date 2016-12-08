package riha_test

import (
	"encoding/json"
	"io"
	"io/ioutil"

	. "github.com/e-gov/riha"

	"net/http"
	"net/http/httptest"

	util "github.com/e-gov/riha/util"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var bufferLength int64 = 1048576

var _ = Describe("RIHA", func() {
	var router *mux.Router
	var recorder *httptest.ResponseRecorder
	var request *http.Request

	BeforeEach(func() {
		util.SetupSvcLogging()

		router = NewRouter()
		recorder = httptest.NewRecorder()

		InitStorage("RIHA_dump.json")

	})
	Describe("List all the systems", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/systems.json", nil)
		})

		Context("Happy day", func() {
			It("Should return 200 and valid systems", func() {
				_ = getSystems(router, recorder, request)
			})

			It("Should return meaningful systems", func() {
				s := getSystems(router, recorder, request)
				Expect(len(s) > 0).To(BeTrue())
				Expect(s[0].Name == "").NotTo(BeTrue())
				Expect(s[0].Shortname == "").NotTo(BeTrue())
				Expect(s[0].Owner).NotTo(BeNil())
			})
		})
	})
})

func getSystems(router *mux.Router, recorder *httptest.ResponseRecorder, request *http.Request) []System {
	var s []System

	router.ServeHTTP(recorder, request)
	Expect(recorder.Code).To(Equal(200))

	body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
	Expect(err).To(BeNil())

	err = json.Unmarshal(body, &s)
	Expect(err).To(BeNil())

	return s
}
