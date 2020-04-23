package notifier

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/notifier/config"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/service"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/store"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

type Server struct {
	config             *config.Config
	context            context.Context
	applicationStore   core.ApplicationStore
	applicationService core.ApplicationService
}

func New(config *config.Config) *Server {
	return &Server{
		config:  config,
		context: context.TODO(),
	}
}

func (s *Server) Start() error {
	db, err := gorm.Open("postgres", s.config.PostgresURL)
	if err != nil {
		return err
	}
	defer db.Close()

	amqpConn, err := amqp.Dial(s.config.RabbitURL)
	if err != nil {
		return err
	}
	defer amqpConn.Close()

	s.applicationStore = store.NewApplicationStore(db)
	s.applicationService = service.NewApplicationService(amqpConn)

	return s.ListenEvents()
}

func (s *Server) Shutdown(ctx context.Context) error {
	//context.WithTimeout(m.context, 15*time.Second)
	return nil
}
