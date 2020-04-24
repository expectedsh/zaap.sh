package deployer

import (
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func (s *Server) ListenEvents() error {
	ch, err := s.amqpConn.Channel()
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare("application_events", amqp.ExchangeTopic, true, false, false, false, nil)
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare("deployer", false, false, false, false, nil)
	if err != nil {
		return err
	}

	err = ch.QueueBind(q.Name, "ApplicationDeploymentRequest", "application_events", false, nil)
	if err != nil {
		return err
	}

	//err = ch.QueueBind(q.Name, "ApplicationDeleted", "application_events", false, nil)
	//if err != nil {
	//	return err
	//}

	err = ch.Qos(1, 0, false)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	for msg := range msgs {
		logrus.Info(string(msg.Body))
	}

	return nil
}

func (s *Server) dispatchEvent(msg proto.Message) error {
	switch event := msg.(type) {
	case *protocol.ApplicationDeploymentRequest:
		logrus.Info(event.ApplicationId, event.DeploymentId)
	}
	return nil
}
