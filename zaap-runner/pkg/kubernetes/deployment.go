package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) DeploymentGet(application *protocol.Application) (*appsv1.Deployment, error) {
	deployment, err :=  c.client.AppsV1().Deployments(c.namespace).Get(c.deploymentName(application), metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return deployment, nil
}

func (c *Client) DeploymentCreate(application *protocol.Application) error {
	_, err := c.client.AppsV1().Deployments(c.namespace).Create(c.toDeployment(application))
	if err != nil {
		return err
	}
	_, err = c.client.CoreV1().Services(c.namespace).Create(c.toService(application))
	return err
}

func (c *Client) DeploymentUpdate(application *protocol.Application) error {
	_, err := c.client.AppsV1().Deployments(c.namespace).Update(c.toDeployment(application))
	return err
}

func (c *Client) DeploymentDelete(application *protocol.Application) error {
	return c.client.AppsV1().Deployments(c.namespace).Delete(c.deploymentName(application), &metav1.DeleteOptions{})
}
