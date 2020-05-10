package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	networkv1 "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
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
	_, err := c.client.NetworkingV1beta1().Ingresses(c.namespace).Create(toIngress(application))
	return err
}

func (c *Client) IngressUpdate(application *runnerpb.Application) error {
	_, err := c.client.NetworkingV1beta1().Ingresses(c.namespace).Update(toIngress(application))
	return err
}

func (c *Client) IngressDelete(name string) error {
	return c.client.NetworkingV1beta1().Ingresses(c.namespace).Delete(name, &metav1.DeleteOptions{})
}

func toIngress(application *runnerpb.Application) *networkv1.Ingress {
	var rules []networkv1.IngressRule

	for _, domain := range application.Domains {
		rules = append(rules, networkv1.IngressRule{
			Host: domain,
			IngressRuleValue: networkv1.IngressRuleValue{
				HTTP: &networkv1.HTTPIngressRuleValue{
					Paths: []networkv1.HTTPIngressPath{
						{
							Path: "/",
							Backend: networkv1.IngressBackend{
								ServiceName: application.Name,
								ServicePort: intstr.FromString("http"),
							},
						},
					},
				},
			},
		})
	}

	meta := toObjectMeta(application)
	meta.Annotations = map[string]string{
		"traefik.ingress.kubernetes.io/router.entrypoints": "web",
	}

	return &networkv1.Ingress{
		ObjectMeta: meta,
		Spec: networkv1.IngressSpec{
			Rules: rules,
		},
	}
}
