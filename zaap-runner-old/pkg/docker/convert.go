package docker

import (
	"fmt"
	"github.com/docker/docker/api/types/swarm"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func ToSwarmSpec(network string, application *protocol.Application) swarm.ServiceSpec {
	replicas := uint64(application.Replicas)

	rawPort := "80"
	if application.Environment["PORT"] != "" {
		rawPort = application.Environment["PORT"]
	}

	port, err := strconv.Atoi(rawPort)
	if err != nil {
		logrus.WithError(err).WithField("port", port).Warn("invalid port")
	}

	var hosts []string

	for _, domain := range application.Domains {
		hosts = append(hosts, fmt.Sprintf("Host(`%v`)", domain))
	}

	rule := strings.Join(hosts, " || ")

	labels := map[string]string{
		"zaap-app-id":        application.Id,
		"zaap-app-name":      application.Name,
		"zaap-deployment-id": application.DeploymentId,
		"traefik.enable":     "true",
		fmt.Sprintf("traefik.http.services.%v.loadbalancer.server.port", application.Name): rawPort,
		fmt.Sprintf("traefik.http.routers.%v-http.rule", application.Name):                 rule,
		fmt.Sprintf("traefik.http.routers.%v-http.entrypoints", application.Name):          "http",
		fmt.Sprintf("traefik.http.routers.%v-https.rule", application.Name):                rule,
		fmt.Sprintf("traefik.http.routers.%v-https.entrypoints", application.Name):         "https",
		fmt.Sprintf("traefik.http.routers.%v-https.tls.certresolver", application.Name):    "simple",
	}

	return swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name:   application.Name + "_" + application.Id,
			Labels: labels,
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: application.Image,
				Env:   ConvertEnv(application.Environment),
			},
		},
		Networks: []swarm.NetworkAttachmentConfig{
			{
				Target: network,
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
