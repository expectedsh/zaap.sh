package kubernetes

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ClusterRoleList() ([]*runnerpb.ClusterRole, error) {
	roles, err := c.client.RbacV1().ClusterRoles().List(metav1.ListOptions{})
	if err != nil {

		return nil, err
	}

	var result []*runnerpb.ClusterRole

	for _, role := range roles.Items {
		result = append(result, &runnerpb.ClusterRole{
			Name: role.Name,
		})
	}

	return result, nil
}
