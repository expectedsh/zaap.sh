package daemon

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"

	application "github.com/remicaumette/zaap.sh/zaap-services/pkg/models/applications"
)

type daemon struct {
	dockerCli      *client.Client
	schedulerToken string
}

func newDaemon(client *client.Client, token string) *daemon {
	return &daemon{
		dockerCli:      client,
		schedulerToken: token,
	}
}

func (d *daemon) deployApplication(payload application.DeploymentPayload) error {
	logrus.WithField("daemon-func", "deployApplication").WithField("payload", payload).Info()
	var replicas = uint64(1)

	_, err := d.dockerCli.ServiceCreate(context.Background(), swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: payload.Application.ID + "_" + payload.Application.Name,
			Labels: map[string]string{
				"container_id": payload.Application.ID,
				"user_id":      payload.Application.UserID,
				"name":         payload.Application.Name,
			},
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: payload.Application.Name,
				Env:   convertEnvironmentVariables(payload.Application.Environment),
			},
		},
		Mode: swarm.ServiceMode{
			Replicated: &swarm.ReplicatedService{
				Replicas: &replicas,
			},
		},
	}, types.ServiceCreateOptions{})

	return err
}

func convertEnvironmentVariables(environment map[string]string) []string {
	var env []string
	for key, value := range environment {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}

	return env
}
