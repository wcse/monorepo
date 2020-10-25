package server

import (
	"context"
	"fmt"
	conf "github.com/wcse/monorepo/backend/api/config/feed"
	"github.com/wcse/monorepo/backend/api/feed"
	"github.com/wcse/monorepo/backend/shared/config"
	"github.com/wcse/monorepo/backend/shared/logging"
	"net"
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Run ...
func Run(f *config.Flags) {
	// Use a temporary logger to parse the configuration and output.
	tmpLogger := logging.NewTmpLogger().With(zap.String("filename", f.ConfigPath))

	var cfg conf.Config
	if err := config.ParseFile(f.ConfigPath, &cfg, f.Template); err != nil {
		tmpLogger.Fatal("parsing configuration failed", zap.Error(err))
	}

	if err := cfg.Validate(); err != nil {
		tmpLogger.Fatal("validating configuration failed", zap.Error(err))
	}

	if f.Validate {
		tmpLogger.Info("configuration validation was successful")
		os.Exit(0)
	}

	RunWithConfig(&cfg)
}

// RunWithConfig ...
func RunWithConfig(cfg *conf.Config) {
	// Init the logger.
	logger, err := logging.NewLogger(cfg.Logger)
	if err != nil {
		logging.NewTmpLogger().Fatal("could not instantiate logger", zap.Error(err))
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			panic(err)
		}
	}()

	// Init listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Listener.GetTcp().Address, cfg.Listener.GetTcp().Port))
	if err != nil {
		logger.Fatal("failed to new listener", zap.Error(err))
	}
	server := grpc.NewServer()
	serverIml := newServerImpl(logger)
	feed.RegisterFeedServer(server, serverIml)
	if err = server.Serve(listener); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
	// Lock by Serve
}

func newServerImpl(logger *zap.Logger) *serverImpl {
	return &serverImpl{
		logger: logger.With(zap.String("serviceName", "feed")),
	}
}

type serverImpl struct {
	feed.UnimplementedFeedServer
	logger *zap.Logger
}

// Register will be called from client-side
func (s *serverImpl) Post(_ context.Context, req *feed.PostRequest) (*feed.PostReply, error) {
	s.logger.Debug("Post", zap.Any("req", req))

	return &feed.PostReply{}, nil
}
