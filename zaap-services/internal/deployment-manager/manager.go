package deployment_manager

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/deployment-manager/config"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/service"
	"github.com/streadway/amqp"
)

type Manager struct {
	context            context.Context
	config             *config.Config
	applicationService core.ApplicationService
	errors             chan error
}

func New(config *config.Config) *Manager {
	return &Manager{
		context: context.TODO(),
		config:  config,
		errors:  make(chan error),
	}
}

func (m *Manager) Start() error {
	amqpConn, err := amqp.Dial(m.config.RabbitURL)
	if err != nil {
		return err
	}
	defer amqpConn.Close()

	m.applicationService = service.NewApplicationService(amqpConn)

	go m.ApplicationEventsHandler()

	return <-m.errors
}

func (m *Manager) Shutdown(ctx context.Context) error {
	//context.WithTimeout(m.context, 15*time.Second)
	return nil
}
