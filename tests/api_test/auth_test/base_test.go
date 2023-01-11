package auth_test

import (
	"gingonic/route/api"
	"gingonic/tests/api_test"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := gin.New()
	api.RegisterAPI(router)

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
