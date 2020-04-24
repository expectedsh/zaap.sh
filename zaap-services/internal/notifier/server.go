package notifier

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/notifier/config"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/notifier/notifiers"
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
	notifier           notifiers.Notifier
	applicationStore   core.ApplicationStore
	applicationService core.ApplicationService
}

func New(config *config.Config) *Server {
	return &Server{
		config:  config,
		context: context.TODO(),
		notifier: notifiers.NewDiscordNotifier("https://discordapp.com/api/webhooks/702952793115459764/TQo8AVcYaTiEUJVv1Zg0gfmImHacjv6ciAlOBKvrs3F0Sb8HAJNmqZyKGmwGn6c264g5"),
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

	queueConfig := messaging.NewSimpleWorkingQueue(service.ApplicationEventsExchange, "notifier")
	subscriber := messaging.NewSubscriber(amqpConn, queueConfig)
	subscriber.RegisterHandler(s.HandleApplicationDeploymentRequested)
	subscriber.RegisterHandler(s.HandleApplicationDeleted)

	return subscriber.Subscribe(s.context)
}

func (s *Server) Shutdown(ctx context.Context) error {
	//context.WithTimeout(m.context, 15*time.Second)
	return nil
}
