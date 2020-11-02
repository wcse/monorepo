package server

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net"
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/wcse/monorepo/backend/api/comment"
	conf "github.com/wcse/monorepo/backend/api/config/comment"
	"github.com/wcse/monorepo/backend/comment/ent"
	"github.com/wcse/monorepo/backend/shared/config"
	"github.com/wcse/monorepo/backend/shared/logging"
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

	// Init database ent
	dbEnt, err := ent.Open("mysql", cfg.Database.Url)
	if err != nil {
		logger.Fatal("failed to connect to mysql", zap.Error(err))
	}
	defer func() {
		if err := dbEnt.Close(); err != nil {
			panic(err)
		}
	}()
	if err := dbEnt.Schema.Create(context.Background()); err != nil {
		logger.Fatal("failed creating schema resources", zap.Error(err))
	}

	// Init listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Listener.GetTcp().Address, cfg.Listener.GetTcp().Port))
	if err != nil {
		logger.Fatal("failed to new listener", zap.Error(err))
	}
	server := grpc.NewServer()
	serverIml := newServerImpl(logger, dbEnt)
	comment.RegisterCommentServer(server, serverIml)
	if err = server.Serve(listener); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
	// Lock by Serve
}

func newServerImpl(logger *zap.Logger, client *ent.Client) *serverImpl {
	return &serverImpl{
		logger: logger.With(zap.String("serviceName", "comment")),
		client: client,
	}
}

type serverImpl struct {
	comment.UnimplementedCommentServer
	logger *zap.Logger
	client *ent.Client
}

func (s *serverImpl) Write(ctx context.Context, req *comment.CommentRequest) (*comment.CommentReply, error) {
	s.logger.Debug("Comment", zap.Any("request", req))

	if err := req.Validate(); err != nil {
		s.logger.Error("bad request", zap.Error(err))
		return nil, err
	}

	// TODO fake userID
	userID := uuid.New().String()

	f, err := s.client.Comment.Create().
		SetFeedID(req.FeedId).
		SetContent(req.Content).
		SetOwnerID(userID).
		Save(ctx)
	if err != nil {
		s.logger.Error("can not save comment", zap.Error(err))
		return nil, err
	}

	return &comment.CommentReply{
		CommentId: f.CommentID.String(),
	}, nil
}
