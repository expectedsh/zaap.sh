package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) DeploymentGet(application *protocol.Application) (*appsv1.Deployment, error) {
	deployment, err := c.client.AppsV1().Deployments(c.namespace).Get(application.Id, metav1.GetOptions{})
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
	if err != nil {
		return err
	}
	_, err = c.client.NetworkingV1beta1().Ingresses(c.namespace).Create(c.toIngress(application))
	return err
}

func (c *Client) DeploymentUpdate(application *protocol.Application) error {
	_, err := c.client.AppsV1().Deployments(c.namespace).Update(c.toDeployment(application))
	if err != nil {
		return err
	}

	svc, err := c.client.CoreV1().Services(c.namespace).Get(application.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if svc != nil {
		service := c.toService(application)
		service.ObjectMeta.ResourceVersion = svc.ObjectMeta.ResourceVersion
		service.Spec.ClusterIP = svc.Spec.ClusterIP
		_, err = c.client.CoreV1().Services(c.namespace).Update(service)
	} else {
		_, err = c.client.CoreV1().Services(c.namespace).Create(c.toService(application))
	}

	if err != nil {
		return err
	}

	_, err = c.client.NetworkingV1beta1().Ingresses(c.namespace).Update(c.toIngress(application))
	return err
}

func (c *Client) DeploymentDelete(application *protocol.Application) error {
	_ = c.client.NetworkingV1beta1().Ingresses(c.namespace).Delete(application.Id, &metav1.DeleteOptions{})
	_ = c.client.CoreV1().Services(c.namespace).Delete(application.Id, &metav1.DeleteOptions{})
	return c.client.AppsV1().Deployments(c.namespace).Delete(application.Id, &metav1.DeleteOptions{})
}
