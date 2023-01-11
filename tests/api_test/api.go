package api_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ApiScenario struct {
	Router  *gin.Engine
	T       *testing.T

	Name           string
	Method         string
	Url            string
	Body           io.Reader
	RequestHeaders map[string]string

	// expectations
	ExpectedStatus     int
	ExpectedContent    string
	NotExpectedContent []string
}

func (a *ApiScenario) Test() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(a.Method, a.Url, a.Body)

	// set default header
	req.Header.Set("Content-Type", gin.MIMEJSON)

	// set scenario headers
	for k, v := range a.RequestHeaders {
		req.Header.Set(k, v)
	}

	a.Router.ServeHTTP(w, req)

	assert.Equal(a.T, a.ExpectedStatus, w.Code)
	assert.Equal(a.T, a.ExpectedContent, w.Body.String())
}
