package deployment_manager

import (
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
)

func (m *Manager) ApplicationEventsHandler() {
	listener, err := m.applicationService.Events(m.context)
	if err != nil {
		m.errors <- err
		return
	}

	for {
		select {
		case err := <-listener.Errors:
			m.errors <- err
			return
		case msg := <-listener.Messages:
			if err := m.dispatchApplicationEvent(msg); err != nil {
				m.errors <- err
			}
		}
	}
}

func (m *Manager) dispatchApplicationEvent(msg proto.Message) error {
	switch payload := msg.(type) {
	case *protocol.ApplicationDeleted:
		logrus.Info(payload)
	default:
		logrus.Warn("unhandled message")
	}
	return nil
}
