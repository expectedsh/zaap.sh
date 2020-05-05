package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
)

func (c *Client) deploymentName(application *protocol.Application) string {
	return application.Name
}

func (c *Client) toObjectMeta(application *protocol.Application) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name: c.deploymentName(application),
		Labels: map[string]string{
			"zaap-application-id":   application.Id,
			"zaap-application-name": application.Name,
			"zaap-deployment-id":    application.DeploymentId,
		},
	}
}

func (c *Client) toDeployment(application *protocol.Application) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: c.toObjectMeta(application),
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(int32(application.Replicas)),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": c.deploymentName(application),
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": c.deploymentName(application),
						"app-id": application.Id,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  application.Name,
							Image: application.Image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80, // todo: use "PORT" env or 80
								},
							},
						},
					},
				},
			},
		},
	}
}

func (c *Client) toService(application *protocol.Application) *apiv1.Service {
	return &apiv1.Service{
		ObjectMeta: c.toObjectMeta(application),
		Spec: apiv1.ServiceSpec{
			Type: apiv1.ServiceTypeClusterIP,
			Ports: []apiv1.ServicePort{
				{
					Port:       80,
					TargetPort: intstr.FromString("http"),
					Protocol:   apiv1.ProtocolTCP,
				},
			},
			Selector: map[string]string{
				"app": c.deploymentName(application),
			},
		},
	}
}
