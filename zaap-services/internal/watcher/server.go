package watcher

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/watcher/config"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/messaging"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/service"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/store"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"sync"
)

type Server struct {
	config             *config.Config
	context            context.Context
	runnerStore        core.RunnerStore
	runnerService      core.RunnerService
	applicationStore   core.ApplicationStore
	applicationService core.ApplicationService
	watcherMutex       sync.Mutex
	watchers           map[uuid.UUID]*Watcher
}

func New(config *config.Config) *Server {
	return &Server{
		config:   config,
		context:  context.TODO(),
		watchers: map[uuid.UUID]*Watcher{},
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

	s.runnerStore = store.NewRunnerStore(db)
	s.runnerService = service.NewRunnerService(amqpConn)
	s.applicationStore = store.NewApplicationStore(db)
	s.applicationService = service.NewApplicationService(amqpConn)

	queueConfig := messaging.NewSimpleWorkingQueue(service.ApplicationEventsExchange.Name(), "watcher")
	subscriber := messaging.NewSubscriber(amqpConn, service.ApplicationEventsExchange, queueConfig)

	runners, err := s.runnerStore.List(s.context)
	if err != nil {
		return err
	}
	for _, runner := range *runners {
		go s.watchRunner(runner)
	}

	return subscriber.Subscribe(s.context)
}

func (s *Server) Shutdown(ctx context.Context) error {
	//context.WithTimeout(m.context, 15*time.Second)
	return nil
}
