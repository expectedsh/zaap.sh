package watcher

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"github.com/sirupsen/logrus"
	"time"
)

type Watcher struct {
	logger *logrus.Entry
	server *Server
	runner *core.Runner
	ctx    context.Context
}

func (w *Watcher) Register() {
	w.server.watcherMutex.Lock()
	w.server.watchers[w.runner.ID] = w
	w.server.watcherMutex.Unlock()
}

func (w *Watcher) Unregister() {
	w.server.watcherMutex.Lock()
	delete(w.server.watchers, w.runner.ID)
	w.server.watcherMutex.Unlock()
}

func (w *Watcher) WatchEvents() error {
	client, conn, err := w.server.runnerService.NewConnection(w.runner)
	if err != nil {
		return err
	}
	defer conn.Close()

	events, err := client.Events(w.ctx, &protocol.GetEventsRequest{})
	if err != nil {
		return err
	}

	for {
		message, err := events.Recv()
		if err != nil {
			return err
		}
		switch message.Event.(type) {
		case *protocol.GetEventsResponse_ApplicationEvent:
			event := message.Event.(*protocol.GetEventsResponse_ApplicationEvent).ApplicationEvent
			logrus.WithField("reason", event.Reason).WithField("message", event.Message).Info("event received")
		}
	}
}

func (s *Server) watchRunner(runner core.Runner) {
	watcher := Watcher{
		logger: logrus.WithField("runner-id", runner.ID).WithField("user-id", runner.UserID.String()),
		server: s,
		runner: &runner,
		ctx:    context.TODO(),
	}
	watcher.Register()
	defer watcher.Unregister()

	closed := false
	go func() {
		<-s.context.Done()
		closed = true
	}()

	for !closed {
		runner.Status = core.RunnerStatusOnline
		if err := s.runnerStore.Update(s.context, &runner); err != nil {
			watcher.logger.WithError(err).Error("failed to update runner status")
		} else if err := watcher.WatchEvents(); err != nil {
			watcher.logger.WithError(err).Error("failed to watch runner events")
			runner.Status = core.RunnerStatusOffline
			if err = s.runnerStore.Update(s.context, &runner); err != nil {
				watcher.logger.WithError(err).Error("failed to update runner status")
			}
		}
		time.Sleep(time.Second * 15)
	}
}
