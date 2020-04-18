package apiserver

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/config"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/handler"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
)

type Server struct {
	config     *config.Config
	httpServer *http.Server
	db         *gorm.DB
}

func New(config *config.Config) *Server {
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
	// todo: use real migration
	if err := db.AutoMigrate(&core.User{}, &core.Application{}, &core.Deployment{}).Error; err != nil {
		return err
	}

	s.httpServer = &http.Server{
		Addr:    s.config.Addr,
		Handler: handler.New(s.config, s.db),
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.httpServer.SetKeepAlivesEnabled(false)
	return s.httpServer.Shutdown(ctx)
}
