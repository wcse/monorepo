package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	cupload "github.com/wcse/monorepo/backend/api/config/upload"
	"github.com/wcse/monorepo/backend/shared/config"
	"github.com/wcse/monorepo/backend/shared/logging"
)

// Response ...
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

// Run ...
func Run(f *config.Flags) {
	// Use a temporary logger to parse the configuration and output.
	tmpLogger := logging.NewTmpLogger().With(zap.String("filename", f.ConfigPath))

	var cfg cupload.Config
	if err := config.ParseFile(f.ConfigPath, &cfg, f.Template); err != nil {
		tmpLogger.Fatal("parsing configuration failed", zap.Error(err))
	}

	RunWithConfig(&cfg)
}

// RunWithConfig ...
func RunWithConfig(cfg *cupload.Config) {
	// Init the logger
	logger, err := logging.NewLogger(cfg.Logger)
	if err != nil {
		logging.NewTmpLogger().Fatal("could not instantiate logger", zap.Error(err))
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			panic(err)
		}
	}()

	// init server
	server := newUploadServer(logger, cfg)
	logger.Info("bind handle function for paths `/`, `/upload`")
	// bind handler
	http.HandleFunc("/", server.status)
	http.HandleFunc("/upload", server.upload)

	// listener
	address := fmt.Sprintf("%s:%d", cfg.Listener.GetTcp().Address, cfg.Listener.GetTcp().Port)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}

func newUploadServer(logger *zap.Logger, cfg *cupload.Config) *server {
	return &server{
		logger: logger.With(zap.String("serviceName", "upload")),
		cfg:    cfg,
	}
}

// UploadServer ...
type UploadServer interface {
	status(w http.ResponseWriter, r *http.Request)
	upload(w http.ResponseWriter, r *http.Request)
}

type server struct {
	logger *zap.Logger
	cfg    *cupload.Config
}

func (s *server) status(w http.ResponseWriter, r *http.Request) {
	s.logger.Info(fmt.Sprintf("health check status: Method: %s", r.Method))
	// create response
	data := make(map[string]string)
	data["message"] = "OK"
	response := Response{
		Success: true,
		Data:    data,
	}

	// make json
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set content-type and return
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
