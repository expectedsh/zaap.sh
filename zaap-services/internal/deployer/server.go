package deployer

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/deployer/config"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/messaging"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/service"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/store"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

type Server struct {
	config             *config.Config
	context            context.Context
	amqpConn           *amqp.Connection
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
	s.amqpConn = amqpConn

	s.applicationStore = store.NewApplicationStore(db)
	s.applicationService = service.NewApplicationService(amqpConn)

	queueConfig := messaging.NewSimpleWorkingQueue(service.ApplicationEventsExchange, "deployer")
	subscriber := messaging.NewSubscriber(amqpConn, queueConfig)
	subscriber.RegisterHandler(s.DeploymentHandler)

	return subscriber.Subscribe(s.context)
}

func (s *Server) Shutdown(ctx context.Context) error {
	//context.WithTimeout(m.context, 15*time.Second)
	return nil
}
