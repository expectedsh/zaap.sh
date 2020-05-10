package notifier

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/notifier/notifiers"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/connector/postgres"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/connector/rabbitmq"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/messaging"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/service"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/store"
)

type Server struct {
	context            context.Context
	notifier           notifiers.Notifier
	applicationStore   core.ApplicationStore
	applicationService core.ApplicationService
}

func New() *Server {
	return &Server{
		context:  context.TODO(),
		notifier: notifiers.NewDiscordNotifier("https://discordapp.com/api/webhooks/702952793115459764/TQo8AVcYaTiEUJVv1Zg0gfmImHacjv6ciAlOBKvrs3F0Sb8HAJNmqZyKGmwGn6c264g5"),
	}
}

func (s *Server) Start() error {
	db, err := postgres.Connect(nil)
	if err != nil {
		return err
	}
	defer db.Close()

	amqpConn, err := rabbitmq.Connect(nil)
	if err != nil {
		return err
	}
	defer amqpConn.Close()

	s.applicationStore = store.NewApplicationStore(db)
	s.applicationService = service.NewApplicationService(amqpConn)

	queueConfig := messaging.NewSimpleWorkingQueue(service.ApplicationEventsExchange.Name(), "notifier")
	subscriber := messaging.NewSubscriber(amqpConn, service.ApplicationEventsExchange, queueConfig)
	subscriber.RegisterHandler(s.HandleApplicationDeploymentRequested)
	subscriber.RegisterHandler(s.HandleApplicationDeleted)

	return subscriber.Subscribe(s.context)
}

func (s *Server) Shutdown(ctx context.Context) error {
	//context.WithTimeout(m.context, 15*time.Second)
	return nil
}
