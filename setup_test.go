package main

import (
	// "net/http"
	// "net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	// set Gin to test mode
	gin.SetMode(gin.TestMode)

	// run the other test functions
	os.Exit(m.Run())
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context)  {
		c.String(200, "pong")
	})
	return r
}

// func getRouter(withTemplates bool) *gin.Engine {
// 	r := gin.Default()
// 	if withTemplates {
// 		r.LoadHTMLGlob("templates/*")
// 	}
// 	return r
// }

// // helper function to process a request and test its response. true if the test is successful
// func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
// 	// create a response recorder
// 	w := httptest.NewRecorder()

// 	// Create the service and process the above request
// 	r.ServeHTTP(w, req)

// 	// if there is no response recorder, fail the test
// 	if !f(w) {
// 		t.Fail()
// 	}
// }