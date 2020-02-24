package server

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"

	"github.com/philips-labs/garo/rpc"
	"github.com/philips-labs/garo/rpc/garo"
)

// Config configures the server
type Config struct {
	Addr      string
	TLSConfig *tls.Config
}

// Run sets up and starts a TLS server that can be cancelled usting the
// given configuration.
func Run(ctx context.Context, conf Config) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", conf.Addr)
	if err != nil {
		return err
	}
	var listener net.Listener
	listener, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	if conf.TLSConfig != nil {
		listener = tls.NewListener(listener, conf.TLSConfig)
	}

	server := &rpc.Server{}
	twirpHandler := garo.NewAgentConfigurationServiceServer(server, nil)

	srv := http.Server{
		Addr:    conf.Addr,
		Handler: twirpHandler,
	}

	err = srv.Serve(listener)
	return err
}
