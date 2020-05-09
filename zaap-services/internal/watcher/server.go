package watcher

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/watcher/config"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/service"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/store"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
	"time"
)

type Server struct {
	config             *config.Config
	context            context.Context
	done               bool
	runnerStore        core.RunnerStore
	runnerService      core.RunnerService
	applicationStore   core.ApplicationStore
	applicationService core.ApplicationService
}

func New(config *config.Config) *Server {
	return &Server{
		config:  config,
		done:    false,
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

	s.runnerStore = store.NewRunnerStore(db)
	s.runnerService = service.NewRunnerService(amqpConn)
	s.applicationStore = store.NewApplicationStore(db)
	s.applicationService = service.NewApplicationService(amqpConn)

	for !s.done {
		runners, err := s.runnerStore.List(s.context)
		if err != nil {
			return err
		}

		for _, runner := range *runners {
			go s.updateRunner(runner)
		}

		time.Sleep(time.Second * 15)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.done = true
	s.context = ctx
	return nil
}
