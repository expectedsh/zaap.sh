package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
)

func (c *Client) labels(application *protocol.Application) map[string]string {
	return map[string]string{
		"zaap-application-id":   application.Name,
		"zaap-application-name": application.Name,
		"zaap-deployment-id":    application.DeploymentId,
	}
}

func (c *Client) toObjectMeta(application *protocol.Application) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:   application.Name,
		Labels: c.labels(application),
	}
}

func (c *Client) toDeployment(application *protocol.Application) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: c.toObjectMeta(application),
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(int32(application.Replicas)),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"zaap-application-name": application.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: c.labels(application),
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
				"zaap-application-name": application.Name,
			},
		},
	}
}

func (c *Client) toIngress(application *protocol.Application) *networkv1.Ingress {
	var rules []networkv1.IngressRule

	for _, domain := range application.Domains {
		rules = append(rules, networkv1.IngressRule{
			Host: domain,
		})
	}

	meta := c.toObjectMeta(application)
	meta.Annotations = map[string]string{
		"traefik.ingress.kubernetes.io/router.entrypoints": "web,websecure",
	}

	return &networkv1.Ingress{
		ObjectMeta: meta,
		Spec: networkv1.IngressSpec{
			Rules: rules,
			Backend: &networkv1.IngressBackend{
				ServiceName: application.Name,
				ServicePort: intstr.FromString("http"),
			},
		},
	}
}
