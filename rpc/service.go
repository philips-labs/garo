package rpc

import (
	"context"
	"fmt"

	"github.com/twitchtv/twirp"

	"github.com/philips-labs/garo/rpc/garo"
)

// Service implements the Garo service
type Service struct{}

// GetRepositoryConfiguration returns the action runner configuration for a given repository
func (s *Service) GetRepositoryConfiguration(ctx context.Context, r *garo.GetRepoConfigurationRequest) (*garo.RepoConfigurationResponse, error) {
	if r.Organisation == "" {
		return nil, twirp.InvalidArgumentError("organisation", "You need to mandatory provide a organization")
	}
	if r.Repository == "" {
		return nil, twirp.InvalidArgumentError("repository", "You need to mandatory provide a repository")
	}

	// TODO: fetch configuration for repository
	return &garo.RepoConfigurationResponse{
		Repository:           fmt.Sprintf("%s/%s", r.Organisation, r.Repository),
		MaxConcurrentRunners: 1,
	}, nil
}
