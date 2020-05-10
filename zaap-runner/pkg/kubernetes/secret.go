package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) SecretImagePullList() ([]*runnerpb.ImagePullSecret, error) {
	roles, err := c.client.CoreV1().Secrets(c.namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var result []*runnerpb.ImagePullSecret

	for _, role := range roles.Items {
		if role.Type == corev1.SecretTypeDockerConfigJson {
			result = append(result, &runnerpb.ImagePullSecret{
				Name: role.Name,
			})
		}
	}

	return result, nil
}
