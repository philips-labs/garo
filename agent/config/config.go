package config

import (
	"errors"
)

var (
	// ErrNoConfig thrown when this agent has no config for this repository
	ErrNoConfig = errors.New("No config for this repository")
)

// RepoConfigurations a map of Configurations per repository
type RepoConfigurations map[string]RepoConfig

// RepoConfig holds configurations for a repository
type RepoConfig struct {
	MaxConcurrentRunners uint32
}
