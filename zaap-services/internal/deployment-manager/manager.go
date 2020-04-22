package deployment_manager

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/deployment-manager/config"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/service"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Manager struct {
	context context.Context
	config  *config.Config
}

func New(config *config.Config) *Manager {
	return &Manager{
		context: context.TODO(),
		config:  config,
	}
}

func (m *Manager) Start() error {
	amqpConn, err := amqp.Dial(m.config.RabbitURL)
	if err != nil {
		return err
	}
	defer amqpConn.Close()

	applicationService := service.NewApplicationService(amqpConn)

	listener, err := applicationService.Events(m.context)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-listener.Errors:
			return err
		case msg := <-listener.Messages:
			m.HandleMessage(msg)
		}
	}

	return nil
}

func (m *Manager) HandleMessage(msg proto.Message) error {
	switch payload := msg.(type) {
	case *protocol.ApplicationDeleted:
		logrus.Info(payload)
	default:
		logrus.Warn("unhandled message")
	}
	return nil
}

func (m *Manager) Shutdown(ctx context.Context) error {
	//context.WithTimeout(m.context, 15*time.Second)
	return nil
}
