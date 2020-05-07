package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) DeploymentCreateOrUpdate(application *protocol.Application) error {
	_, err := c.client.AppsV1().Deployments(c.namespace).Get(application.Id, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return c.DeploymentCreate(application)
	} else if err != nil {
		return err
	}
	return c.DeploymentUpdate(application)
}

func (c *Client) DeploymentCreate(application *protocol.Application) error {
	_, err := c.client.AppsV1().Deployments(c.namespace).Create(c.toDeployment(application))
	return err
}

func (c *Client) DeploymentUpdate(application *protocol.Application) error {
	_, err := c.client.AppsV1().Deployments(c.namespace).Update(c.toDeployment(application))
	return err
}

func (c *Client) DeploymentDelete(name string) error {
	return c.client.AppsV1().Deployments(c.namespace).Delete(name, &metav1.DeleteOptions{})
}
