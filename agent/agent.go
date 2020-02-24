package agent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/philips-labs/garo/rpc/garo"
)

// Config configures the agent
type Config struct {
	ServerAddr string
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

	fmt.Printf("Agent configuration: %+v\n", cfg)

	return nil
}

func initClient(addr string) garo.AgentConfigurationService {
	return garo.NewAgentConfigurationServiceProtobufClient(addr, &http.Client{})
}
