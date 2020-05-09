package kubernetes

import (
	"fmt"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (c *Client) GetStatus(applicationId, deploymentId string) (protocol.ApplicationStatus, error) {
	rsl, err := c.client.AppsV1().ReplicaSets(c.namespace).List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("zaap-application-id=%v,zaap-deployment-id=%v", applicationId, deploymentId),
	})
	if err != nil {
		return protocol.ApplicationStatus_UNKNOWN, err
	}

	if len(rsl.Items) != 1 {
		return protocol.ApplicationStatus_UNKNOWN, nil
	}

	rs := rsl.Items[0]
	hash := rs.ObjectMeta.Labels["pod-template-hash"]

	podsl, err := c.client.CoreV1().Pods(c.namespace).List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("pod-template-hash=%v", hash),
	})
	if err != nil {
		return protocol.ApplicationStatus_UNKNOWN, err
	}
	pods := podsl.Items

	podEvents, err := c.getPodsLastEvent(pods)
	if err != nil {
		return protocol.ApplicationStatus_UNKNOWN, err
	}

	pending, running, failed := 0, 0, 0

	for _, pod := range pods {
		switch pod.Status.Phase {
		case corev1.PodPending:
			if event := podEvents[pod.UID]; event != nil && event.Reason == "Failed" {
				failed++
			} else {
				pending++
			}
		case corev1.PodRunning:
			running++
		case corev1.PodFailed:
			failed++
		}
	}

	if pending > running && pending > failed {
		return protocol.ApplicationStatus_DEPLOYING, nil
	} else if running > pending && running > failed {
		return protocol.ApplicationStatus_RUNNING, nil
	} else if failed > pending && failed > running {
		return protocol.ApplicationStatus_FAILED, nil
	}

	return protocol.ApplicationStatus_RUNNING, nil
}

func (c *Client) getPodsLastEvent(pods []corev1.Pod) (map[types.UID]*corev1.Event, error) {
	events, err := c.client.CoreV1().Events(c.namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	result := make(map[types.UID]*corev1.Event)

	for _, event := range events.Items {
		for _, pod := range pods {
			if event.InvolvedObject.UID == pod.UID {
				curr := result[pod.UID]
				if curr == nil || event.EventTime.After(curr.EventTime.Time) {
					result[pod.UID] = &event
				}
				break
			}
		}
	}

	return result, nil
}
