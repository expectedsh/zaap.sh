package docker

import (
	"github.com/docker/docker/api/types/swarm"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
)

func ToSwarmSpec(application *protocol.Application) swarm.ServiceSpec {
	replicas := uint64(application.Replicas)
	return swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: application.Name + "_" + application.Id,
			Labels: map[string]string{
				"zaap-app-id":   application.Id,
				"zaap-app-name": application.Name,
			},
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: application.Image,
				Env:   ConvertEnv(application.Environment),
			},
		},
		Mode: swarm.ServiceMode{
			Replicated: &swarm.ReplicatedService{
				Replicas: &replicas,
			},
		},
	}
}

func ConvertEnv(env map[string]string) []string {
	var converted []string
	for k, v := range env {
		converted = append(converted, k+"="+v)
	}
	return converted
}
