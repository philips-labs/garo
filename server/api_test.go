package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"github.com/philips-labs/garo/rpc"
	"github.com/philips-labs/garo/rpc/garo"
)

type registeredRoute struct {
	method string
	route  string
}

func bootstrapAPI() *chi.Mux {
	svc := &rpc.Service{}
	twirpServer := garo.NewAgentConfigurationServiceServer(svc, nil)
	return configureAPI(twirpServer, zap.NewNop())
}

func TestRoutes(t *testing.T) {
	assert := assert.New(t)
	expectedRoutes := []registeredRoute{
		{http.MethodGet, "/"},
		{http.MethodGet, "/ping"},
		{http.MethodDelete, "/twirp/philips.garo.garo.AgentConfigurationService/*"},
		{http.MethodPost, "/twirp/philips.garo.garo.AgentConfigurationService/*"},
		{http.MethodHead, "/twirp/philips.garo.garo.AgentConfigurationService/*"},
		{http.MethodPatch, "/twirp/philips.garo.garo.AgentConfigurationService/*"},
		{http.MethodOptions, "/twirp/philips.garo.garo.AgentConfigurationService/*"},
		{http.MethodPut, "/twirp/philips.garo.garo.AgentConfigurationService/*"},
		{http.MethodGet, "/twirp/philips.garo.garo.AgentConfigurationService/*"},
		{http.MethodTrace, "/twirp/philips.garo.garo.AgentConfigurationService/*"},
		{http.MethodConnect, "/twirp/philips.garo.garo.AgentConfigurationService/*"},
	}

	router := bootstrapAPI()

	routes := make([]registeredRoute, 0)
	err := chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		routes = append(routes, registeredRoute{method, route})
		return nil
	})

	assert.NoError(err, "Failed to walk handlers")
	assert.ElementsMatch(expectedRoutes, routes)
}

func TestGetRoot(t *testing.T) {
	assert := assert.New(t)
	router := bootstrapAPI()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(err, "Failed to create request")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(http.StatusOK, rr.Code, "Invalid status code")
	assert.Equal("", rr.Body.String(), "Invalid response text")
}

func TestGetPing(t *testing.T) {
	assert := assert.New(t)
	router := bootstrapAPI()

	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	assert.NoError(err, "Failed to create request")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(http.StatusOK, rr.Code, "Invalid status code")
	assert.Equal("pong\n", rr.Body.String(), "Invalid response text")
}
