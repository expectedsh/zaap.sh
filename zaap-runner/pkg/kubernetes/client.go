package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Client struct {
	client    kubernetes.Interface
	namespace string
}

func NewClient(namespace string, config *rest.Config) (*Client, error) {
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	_, err = client.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		_, err = client.CoreV1().Namespaces().Create(&apiv1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace,
			},
		})
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return &Client{
		namespace: namespace,
		client:    client,
	}, nil
}

func toLabels(application *runnerpb.Application) map[string]string {
	return map[string]string{
		"zaap-application-id":   application.Id,
		"zaap-application-name": application.Name,
		"zaap-deployment-id":    application.DeploymentId,
	}
}

func toObjectMeta(application *runnerpb.Application) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:   application.Name,
		Labels: toLabels(application),
	}
}
