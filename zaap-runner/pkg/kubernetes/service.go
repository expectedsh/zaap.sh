package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ServiceCreateOrUpdate(application *protocol.Application) error {
	service, err := c.client.CoreV1().Services(c.namespace).Get(application.Name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return c.ServiceCreate(application)
	} else if err != nil {
		return err
	}
	return c.ServiceUpdate(application, service)
}

func (c *Client) ServiceCreate(application *protocol.Application) error {
	_, err := c.client.CoreV1().Services(c.namespace).Create(c.toService(application))
	return err
}

func (c *Client) ServiceUpdate(application *protocol.Application, current *corev1.Service) error {
	service := c.toService(application)

	service.ObjectMeta.ResourceVersion = current.ObjectMeta.ResourceVersion
	service.Spec.ClusterIP = current.Spec.ClusterIP

	_, err := c.client.CoreV1().Services(c.namespace).Update(service)
	return err
}

func (c *Client) ServiceDelete(name string) error {
	return c.client.CoreV1().Services(c.namespace).Delete(name, &metav1.DeleteOptions{})
}
