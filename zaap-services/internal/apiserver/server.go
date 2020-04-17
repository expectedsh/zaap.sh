package apiserver

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/handler"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Server struct {
	config     *Config
	httpServer *http.Server
	db         *gorm.DB
}

func New(config *Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start() error {
	db, err := gorm.Open("postgres", s.config.PostgresURL)
	if err != nil {
		return err
	}
	defer db.Close()
	s.db = db

	s.httpServer = &http.Server{
		Addr:    s.config.Addr,
		Handler: handler.New(db),
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.httpServer.SetKeepAlivesEnabled(false)
	return s.httpServer.Shutdown(ctx)
}
