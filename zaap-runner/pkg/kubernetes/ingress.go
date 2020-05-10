package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) IngressCreateOrUpdate(application *runnerpb.Application) error {
	_, err := c.client.NetworkingV1beta1().Ingresses(c.namespace).Get(application.Name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return c.IngressCreate(application)
	} else if err != nil {
		return err
	}
	return c.IngressUpdate(application)
}

func (c *Client) IngressCreate(application *runnerpb.Application) error {
	_, err := c.client.NetworkingV1beta1().Ingresses(c.namespace).Create(c.toIngress(application))
	return err
}

func (c *Client) IngressUpdate(application *runnerpb.Application) error {
	_, err := c.client.NetworkingV1beta1().Ingresses(c.namespace).Update(c.toIngress(application))
	return err
}

func (c *Client) IngressDelete(name string) error {
	return c.client.NetworkingV1beta1().Ingresses(c.namespace).Delete(name, &metav1.DeleteOptions{})
}
