package kubernetes

import (
	"fmt"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/sirupsen/logrus"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ClusterRoleBindingSync(application *runnerpb.Application) error {
	roleBindings, err := c.client.RbacV1().ClusterRoleBindings().List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("zaap-application-id=%v", application.Id),
	})
	if err != nil {
		return err
	}

	var currentRoles []string

	for _, roleBinding := range roleBindings.Items {
		currentRoles = append(currentRoles, roleBinding.RoleRef.Name)
		if !contains(application.Roles, roleBinding.RoleRef.Name) {
			if err := c.ClusterRoleBindingDelete(application.Name, roleBinding.RoleRef.Name); err != nil {
				logrus.
					WithField("application-name", application.Name).
					WithField("role-name", roleBinding.RoleRef.Name).
					WithError(err).
					Error("failed to delete cluster role binding")
			}
		}
	}

	for _, role := range application.Roles {
		if !contains(currentRoles, role) {
			if err := c.ClusterRoleBindingCreate(application, role); err != nil {
				logrus.
					WithField("application-name", application.Name).
					WithField("role-name", role).
					WithError(err).
					Error("failed to create cluster role binding")
			}
		}
	}

	return nil
}

func (c *Client) ClusterRoleBindingCreate(application *runnerpb.Application, role string) error {
	_, err := c.client.RbacV1().ClusterRoleBindings().Create(toClusterRoleBinding(application, role, c.namespace))
	return err
}

func (c *Client) ClusterRoleBindingDeleteAll(applicationId, applicationName string) error {
	roleBindings, err := c.client.RbacV1().ClusterRoleBindings().List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("zaap-application-id=%v", applicationId),
	})
	if err != nil {
		return err
	}

	for _, roleBinding := range roleBindings.Items {
		if err := c.ClusterRoleBindingDelete(applicationName, roleBinding.RoleRef.Name); err != nil {
			logrus.
				WithField("application-name", applicationName).
				WithField("role-name", roleBinding.RoleRef.Name).
				WithError(err).
				Error("failed to delete cluster role binding")
		}
	}

	return nil
}

func (c *Client) ClusterRoleBindingDelete(applicationName, role string) error {
	return c.client.RbacV1().ClusterRoleBindings().Delete(applicationName+"-"+role, &metav1.DeleteOptions{})
}

func toClusterRoleBinding(application *runnerpb.Application, role, namespace string) *rbacv1.ClusterRoleBinding {
	meta := toObjectMeta(application)
	meta.Name = application.Name + "-" + role

	return &rbacv1.ClusterRoleBinding{
		ObjectMeta: meta,
		RoleRef: rbacv1.RoleRef{
			Kind: "ClusterRole",
			Name: role,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      application.Name,
				Namespace: namespace,
			},
		},
	}
}

func contains(slice []string, item string) bool {
	for _, curr := range slice {
		if item == curr {
			return true
		}
	}
	return false
}
