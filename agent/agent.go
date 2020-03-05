package agent

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/philips-labs/garo/agent/config"
	"github.com/philips-labs/garo/rpc/garo"
)

// Config configures the agent
type Config struct {
	ServerAddr      string
	Logger          *zap.Logger
	RefreshInterval time.Duration
	Repositories    []string
}

// Run runs the agent to manage your github action workers
func Run(ctx context.Context, conf Config) error {
	client := initClient(conf.ServerAddr)

	w := config.NewWatcher(conf.Logger, client, conf.Repositories)
	go w.Watch(ctx, conf.RefreshInterval)

	return nil
}

func initClient(addr string) garo.AgentConfigurationService {
	return garo.NewAgentConfigurationServiceProtobufClient(addr, &http.Client{})
}
