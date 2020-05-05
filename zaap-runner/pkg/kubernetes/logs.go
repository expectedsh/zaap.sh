package kubernetes

import (
	"bufio"
	"context"
	"fmt"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"strings"
	"time"
)

type (
	LogEntry struct {
		Time    time.Time
		Pod     string
		Message string
	}

	logWatcher struct {
		parent context.Context
		client *Client
		logs   chan LogEntry
		pods   map[string]context.CancelFunc
	}
)

func (c *Client) Logs(ctx context.Context, application *protocol.Application) (<-chan LogEntry, error) {
	log := logWatcher{
		parent: ctx,
		client: c,
		logs:   make(chan LogEntry),
		pods:   make(map[string]context.CancelFunc),
	}

	go log.listenEvents(ctx, application)

	return log.logs, nil
}

func (w *logWatcher) listenEvents(ctx context.Context, application *protocol.Application) error {
	watcher, err := w.client.client.CoreV1().Pods(w.client.namespace).Watch(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app-id=%v", application.Id),
	})
	if err != nil {
		return err
	}
	defer watcher.Stop()

	defer func() {
		for _, cancel := range w.pods {
			cancel()
		}
		close(w.logs)
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-watcher.ResultChan():
			if event.Object == nil {
				continue
			}
			pod := event.Object.(*corev1.Pod)

			switch event.Type {
			case watch.Added:
				go w.addPod(pod)
			case watch.Modified:
				switch pod.Status.Phase {
				case corev1.PodRunning:
					go w.addPod(pod)
				case corev1.PodSucceeded, corev1.PodFailed:
					if cancel := w.pods[pod.Name]; cancel != nil {
						cancel()
					}
				}
			case watch.Deleted:
				if cancel := w.pods[pod.Name]; cancel != nil {
					cancel()
				}
			}
		}
	}
}

func (w *logWatcher) addPod(pod *corev1.Pod) error {
	ctx, cancel := context.WithCancel(w.parent)
	w.pods[pod.Name] = cancel

	stream, err := w.client.client.CoreV1().Pods(w.client.namespace).GetLogs(pod.Name, &corev1.PodLogOptions{
		Follow:     true,
		Timestamps: true,
	}).Stream()
	if err != nil {
		return err
	}
	defer stream.Close()

	go func() {
		<-ctx.Done()
		_ = stream.Close()
	}()

	sc := bufio.NewScanner(stream)

	for sc.Scan() {
		args := strings.SplitN(sc.Text(), " ", 2)
		t, err := time.Parse(time.RFC3339Nano, args[0])
		if err != nil {
			continue
		}

		w.logs <- LogEntry{
			Time:    t,
			Pod:     pod.Name,
			Message: args[1],
		}
	}

	return nil
}
