package server

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"

	"github.com/philips-labs/garo/rpc"
	"github.com/philips-labs/garo/rpc/garo"
)

// Config configures the server
type Config struct {
	Addr      string
	TLSConfig *tls.Config
	Logger    *zap.Logger
}

// Server holds the http.Server instance
type Server struct {
	*http.Server
	conf Config
}

// New creates a new instance of Server
func New(conf Config) *Server {
	svc := &rpc.Service{}

	twirpServer := garo.NewAgentConfigurationServiceServer(svc, nil)
	api := configureAPI(twirpServer, conf.Logger)

	srv := http.Server{
		Addr:    conf.Addr,
		Handler: api,
	}

	return &Server{&srv, conf}
}

// Run sets up and starts a TLS server that can be cancelled usting the
// given configuration.
func Run(ctx context.Context, conf Config) error {
	srv := New(conf)
	return srv.Run(ctx)
}

// Run sets up and starts a TLS server that can be cancelled usting the
// given configuration.
func (s *Server) Run(ctx context.Context) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", s.conf.Addr)
	if err != nil {
		return err
	}
	var listener net.Listener
	listener, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	if s.conf.TLSConfig != nil {
		listener = tls.NewListener(listener, s.conf.TLSConfig)
	}
	err = s.Serve(listener)
	return err
}

// GracefulShutdown waits for os.Interrupt to gracefully shutdown the webserver
// there is a timeout of 30 seconds before the shutdown is forced.
func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit

	s.conf.Logger.Info("Server is shutting down", zap.String("reason", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.SetKeepAlivesEnabled(false)
	if err := s.Shutdown(ctx); err != nil {
		s.conf.Logger.Fatal("Could not gracefully shutdown the server", zap.Error(err))
	}
	s.conf.Logger.Info("Server stopped")
}
