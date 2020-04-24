package notifier

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/notifier/notifiers"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"github.com/golang/protobuf/proto"
)

func (s *Server) ListenEvents() error {
	listener, err := s.applicationService.Events(s.context, "")
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-listener.Errors:
			return err
		case message := <-listener.Messages:
			if err := s.dispatchEvent(message); err != nil {
				return err
			}
		}
	}
}

func (s *Server) dispatchEvent(msg proto.Message) error {
	notifier := notifiers.NewDiscordNotifier("https://discordapp.com/api/webhooks/702952793115459764/TQo8AVcYaTiEUJVv1Zg0gfmImHacjv6ciAlOBKvrs3F0Sb8HAJNmqZyKGmwGn6c264g5")
	switch event := msg.(type) {
	case *protocol.ApplicationDeleted:
		return notifier.WhenApplicationDeleted(event.Id, event.Name)
	}
	return nil
}
