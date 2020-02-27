package rpc

import (
	"context"

	"github.com/twitchtv/twirp"

	"github.com/philips-labs/garo/rpc/garo"
)

// Service implements the Garo service
type Service struct{}

// GetAgentConfiguration returns the action runner configuration for a given repository
func (s *Service) GetAgentConfiguration(ctx context.Context, r *garo.GetAgentConfigurationRequest) (*garo.AgentConfigurationResponse, error) {
	if r.Organisation == "" {
		return nil, twirp.InvalidArgumentError("organisation", "You need to mandatory provide a organization")
	}
	if r.Repository == "" {
		return nil, twirp.InvalidArgumentError("repository", "You need to mandatory provide a repository")
	}

	// TODO: fetch configuration for repository
	return &garo.AgentConfigurationResponse{
		TotalAllowedRunners: 1,
	}, nil
}
