package runner

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/sirupsen/logrus"
)

func (r *Runner) Events(req *protocol.GetEventsRequest, srv protocol.Runner_EventsServer) error {
	logrus.Info("getting events")

	events, err := r.client.Events(srv.Context())
	if err != nil {
		return err
	}

	for _ = range events {
		//err = srv.Send(&protocol.GetEventsResponse{
		//	Event: &protocol.GetEventsResponse_ApplicationEvent{
		//		ApplicationEvent: &protocol.ApplicationEvent{
		//			Message:
		//		},
		//	},
		//})
		//if err != nil {
		//	return err
		//}
	}

	return nil
}
