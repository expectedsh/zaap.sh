package apiserver

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/config"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/handler"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/connector/postgres"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/connector/rabbitmq"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"net/http"
)

type Server struct {
	config     *config.Config
	httpServer *http.Server
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start() error {
	db, err := postgres.Connect(nil)
	if err != nil {
		return err
	}
	defer db.Close()

	// todo: use real migration
	if err := db.AutoMigrate(
		&core.User{},
		&core.Application{},
		&core.Deployment{},
		&core.Runner{},
	).Error; err != nil {
		return err
	}

	amqpConn, err := rabbitmq.Connect(nil)
	if err != nil {
		return err
	}
	defer amqpConn.Close()

	s.httpServer = &http.Server{
		Addr:    s.config.Addr,
		Handler: handler.New(s.config, db, amqpConn),
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.httpServer.SetKeepAlivesEnabled(false)
	return s.httpServer.Shutdown(ctx)
}
