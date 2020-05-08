package kubernetes

import (
	"context"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EventEntry struct {

}

func (c *Client) Events(ctx context.Context) (<-chan EventEntry, error) {
	watcher, err := c.client.AppsV1().Deployments(c.namespace).Watch(metav1.ListOptions{
		//LabelSelector: "zaap-application-id",
	})
	if err != nil {
		return nil, err
	}

	events := make(chan EventEntry)

	go func() {
		defer func() {
			close(events)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case result := <-watcher.ResultChan():
				if result.Object == nil {
					continue
				}
				deployment := result.Object.(*appsv1.Deployment)

				//deployment.Status.Replicas == deployment.Status.

				// result.Type == ADDED, application deployé
				// result.Type == MODIFIED, application mise a jour
				// result.Type == DELETED, application supprimé

				logrus.WithField("type", result.Type).Info("receiving events")
				for _, status := range deployment.Status.Conditions {
					logrus.
						WithField("status", status.Status).
						WithField("type", status.Type).
						WithField("message", status.Message).
						WithField("last-update-time", status.LastUpdateTime).
						Info("status event")
				}
				//if event.InvolvedObject.Kind == "Deployment" {
				//	logrus.
				//		WithField("resource-kind", event.InvolvedObject.Kind).
				//		WithField("resource-name", event.InvolvedObject.Name).
				//		WithField("reason", event.Reason).
				//		WithField("message", event.Message).
				//		Info("event")
				//}
			}
		}
	}()

	return events, nil
}
