package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ServiceAccountCreateOrUpdate(application *runnerpb.Application) error {
	_, err := c.client.CoreV1().ServiceAccounts(c.namespace).Get(application.Name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		if application.Roles != nil && len(application.Roles) > 0 {
			return c.ServiceAccountCreate(application)
		}
		return nil
	} else if err != nil {
		return err
	}

	if application.Roles == nil || len(application.Roles) == 0 {
		return c.ServiceAccountDelete(application.Name)
	}

	return c.ServiceAccountUpdate(application)
}

func (c *Client) ServiceAccountCreate(application *runnerpb.Application) error {
	_, err := c.client.CoreV1().ServiceAccounts(c.namespace).Create(toServiceAccount(application))
	return err
}

func (c *Client) ServiceAccountUpdate(application *runnerpb.Application) error {
	_, err := c.client.CoreV1().ServiceAccounts(c.namespace).Update(toServiceAccount(application))
	return err
}

func (c *Client) ServiceAccountDelete(name string) error {
	return c.client.CoreV1().ServiceAccounts(c.namespace).Delete(name, &metav1.DeleteOptions{})
}

func toServiceAccount(application *runnerpb.Application) *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		ObjectMeta: toObjectMeta(application),
	}
}
