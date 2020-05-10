package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"strconv"
)

func (c *Client) DeploymentCreateOrUpdate(application *runnerpb.Application) error {
	_, err := c.client.AppsV1().Deployments(c.namespace).Get(application.Name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return c.DeploymentCreate(application)
	} else if err != nil {
		return err
	}
	return c.DeploymentUpdate(application)
}

func (c *Client) DeploymentCreate(application *runnerpb.Application) error {
	_, err := c.client.AppsV1().Deployments(c.namespace).Create(toDeployment(application))
	return err
}

func (c *Client) DeploymentUpdate(application *runnerpb.Application) error {
	_, err := c.client.AppsV1().Deployments(c.namespace).Update(toDeployment(application))
	return err
}

func (c *Client) DeploymentDelete(name string) error {
	return c.client.AppsV1().Deployments(c.namespace).Delete(name, &metav1.DeleteOptions{})
}

func toDeployment(application *runnerpb.Application) *appsv1.Deployment {
	serviceAccount := ""
	port := 80

	if application.Roles != nil && len(application.Roles) > 0 {
		serviceAccount = application.Name
	}

	if value, ok := application.Environment["PORT"]; ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			logrus.WithField("port", value).WithError(err).Warn("failed to convert port to int, fallbacking to 80")
		} else {
			port = i
		}
	}

	var env []corev1.EnvVar

	for key, value := range application.Environment {
		env = append(env, corev1.EnvVar{
			Name:  key,
			Value: value,
		})
	}

	return &appsv1.Deployment{
		ObjectMeta: toObjectMeta(application),
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(int32(application.Replicas)),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"zaap-application-name": application.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: toLabels(application),
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: serviceAccount,
					Containers: []corev1.Container{
						{
							Name:  application.Name,
							Image: application.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: int32(port),
								},
							},
							Env: env,
						},
					},
					ImagePullSecrets: []corev1.LocalObjectReference{
						{
							Name: "regcred", // todo : flex
						},
					},
				},
			},
		},
	}
}
