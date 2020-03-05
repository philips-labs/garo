package config

import (
	"context"
	"reflect"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/philips-labs/garo/rpc/garo"
)

// Watcher watches for configuration changes at garo server
type Watcher struct {
	logger  *zap.Logger
	client  garo.AgentConfigurationService
	configs RepoConfigurations
	m       sync.Mutex
}

// NewWatcher creates a new instance of a configuration watcher
func NewWatcher(l *zap.Logger, client garo.AgentConfigurationService, repos []string) *Watcher {
	configs := make(RepoConfigurations)
	for _, repo := range repos {
		configs[repo] = RepoConfig{}
	}
	return &Watcher{logger: l, client: client, configs: configs}
}

// Watch starts checking your repositories for configuration changes
func (w *Watcher) Watch(ctx context.Context, refreshInterval time.Duration) {
	if refreshInterval < 500*time.Millisecond {
		w.logger.Info("Refresh interval to low, will be set at 500ms", zap.Duration("refreshInterval", refreshInterval))
		refreshInterval = 500 * time.Millisecond
	}
	for {
		select {
		case <-time.After(refreshInterval):
			for repo, currentCfg := range w.configs {
				w.checkConfig(ctx, repo, currentCfg)
			}
		case <-ctx.Done():
			return
		}
	}
}

// GetConfig retrieves the current configuration for a repository
func (w *Watcher) GetConfig(repo string) (RepoConfig, error) {
	w.m.Lock()
	defer w.m.Unlock()
	if cfg, ok := w.configs[repo]; ok {
		return cfg, nil
	}
	return RepoConfig{}, ErrNoConfig
}

func (w *Watcher) checkConfig(ctx context.Context, repo string, currentCfg RepoConfig) {
	repoLogger := w.logger.With(zap.String("repo", repo))
	defer repoLogger.Sync()

	repoSplit := strings.Split(repo, "/")

	repoLogger.Debug("Fetching configuration")
	cfg, err := w.client.GetRepositoryConfiguration(ctx, &garo.GetRepoConfigurationRequest{
		Organisation: repoSplit[0],
		Repository:   repoSplit[1],
	})
	if err != nil {
		repoLogger.Error("Failed to get configuration", zap.Error(err))
		return
	}
	newCfg := RepoConfig{cfg.MaxConcurrentRunners}
	if !reflect.DeepEqual(currentCfg, newCfg) {
		repoLogger.Info("Configuration changed", zap.Any("cfg", cfg))
		w.m.Lock()
		defer w.m.Unlock()
		w.configs[cfg.Repository] = newCfg
	}
}
