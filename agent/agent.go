package agent

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/philips-labs/garo/rpc/garo"
)

// Config configures the agent
type Config struct {
	ServerAddr string
	Logger     *zap.Logger
}

// Run runs the agent to manage your github action workers
func Run(ctx context.Context, conf Config) error {
	client := initClient(conf.ServerAddr)

	cfg, err := client.GetAgentConfiguration(ctx, &garo.GetAgentConfigurationRequest{
		Organisation: "philips-internal",
		Repository:   "fact-service",
	})
	if err != nil {
		return err
	}

	conf.Logger.Info("Agent configuration", zap.Any("cfg", cfg))

	return nil
}

func initClient(addr string) garo.AgentConfigurationService {
	return garo.NewAgentConfigurationServiceProtobufClient(addr, &http.Client{})
}
