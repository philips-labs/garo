package config_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"github.com/philips-labs/garo/agent/config"
	"github.com/philips-labs/garo/rpc"
)

func TestUnwatchedRepository(t *testing.T) {
	w := config.NewWatcher(zap.NewNop(), &rpc.Service{}, []string{"philips-internal/fact-service"})
	cfg, err := w.GetConfig("marcofranssen/whisky-tango")
	assert.EqualError(t, err, config.ErrNoConfig.Error())
	assert.Equal(t, config.RepoConfig{}, cfg)
}

func TestMinimumInterval(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := "philips-internal/fact-service"

	w := config.NewWatcher(zap.NewNop(), &rpc.Service{}, []string{repo})
	go w.Watch(ctx, 5*time.Millisecond)

	<-time.After(5 * time.Millisecond)

	cfg, err := w.GetConfig(repo)
	assert.NoError(t, err)
	assert.Equal(t, config.RepoConfig{}, cfg)

	<-time.After(500 * time.Millisecond)

	cfg, err = w.GetConfig(repo)
	assert.NoError(t, err)
	assert.Equal(t, config.RepoConfig{MaxConcurrentRunners: 1}, cfg)
}

func TestWatchConfig(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := "philips-internal/fact-service"

	w := config.NewWatcher(zap.NewNop(), &rpc.Service{}, []string{repo})
	go w.Watch(ctx, 550*time.Millisecond)

	cfg, err := w.GetConfig(repo)
	assert.NoError(t, err)
	assert.Equal(t, config.RepoConfig{}, cfg)

	<-time.After(555 * time.Millisecond)

	cfg, err = w.GetConfig(repo)
	assert.NoError(t, err)
	assert.Equal(t, config.RepoConfig{MaxConcurrentRunners: 1}, cfg)
}
