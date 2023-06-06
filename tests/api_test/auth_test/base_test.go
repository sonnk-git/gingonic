package auth_test

import (
	"gingonic/route/api"
	"gingonic/tests/api_test"
	"github.com/gin-gonic/gin"
	"io"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := gin.New()
	api.RegisterAPI(router)

	// disable logging routes
	router.Use(gin.LoggerWithWriter(io.Discard))

	apiScenario := []api_test.ApiScenario{
		{
			Router:             router,
			T:                  t,
			Name:               "",
			Method:             "GET",
			Url:                "/api",
			Body:               nil,
			RequestHeaders:     nil,
			ExpectedStatus:     200,
			ExpectedContent:    "{}",
			NotExpectedContent: nil,
		},
	}

	for _, a := range apiScenario {
		a.Test()
	}

}

func TestLogout(t *testing.T) {
	//
}

func TestLogin(t *testing.T) {
	//
}
