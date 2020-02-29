package rpc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/philips-labs/garo/rpc/garo"

	"github.com/philips-labs/garo/rpc"
)

var (
	svc garo.AgentConfigurationService
)

func init() {
	svc = &rpc.Service{}
}

func TestGetRepoConfigurationWithoutOrganization(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := svc.GetRepositoryConfiguration(ctx, &garo.GetRepoConfigurationRequest{
		Repository: "fact-service",
	})
	assert.Error(t, err)
	assert.Nil(t, cfg)
}

func TestGetRepoConfigurationWithoutRepository(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := svc.GetRepositoryConfiguration(ctx, &garo.GetRepoConfigurationRequest{
		Organisation: "philips-internal",
	})
	assert.Error(t, err)
	assert.Nil(t, cfg)
}

func TestGetRepoConfiguration(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := svc.GetRepositoryConfiguration(ctx, &garo.GetRepoConfigurationRequest{
		Organisation: "philips-internal",
		Repository:   "fact-service",
	})
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
}
