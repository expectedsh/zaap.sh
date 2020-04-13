package daemon

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/filters"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/core"
	"net/url"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

type Config struct {
	SchedulerToken string `envconfig:"SCHEDULER_TOKEN"`
	ControllerAddr string `envconfig:"CONTROLLER_ADDR" default:"localhost:9090"`
}

type Daemon struct {
	config *Config
	docker *client.Client
}

func New(config Config) *Daemon {
	return &Daemon{
		config: &config,
	}
}

func (s *Daemon) Start(ctx context.Context) error {
	if s.config.SchedulerToken == "" {
		return errors.New("scheduler token required")
	}
	logrus.Infof("using scheduler token %v", s.config.SchedulerToken)

	docker, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	defer docker.Close()
	s.docker = docker

	controllerUrl := url.URL{Scheme: "ws", Host: s.config.ControllerAddr, Path: "/"}
	return RegisterControllerConsumer(ctx, s.config.SchedulerToken, controllerUrl, s.docker)
}

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

func (d *daemon) getApplication(id string) (*swarm.Service, error) {
	args := filters.NewArgs()
	args.Add("label", fmt.Sprintf("zaap-app-id=%s", id))
	services, err := d.dockerCli.ServiceList(context.Background(), types.ServiceListOptions{
		Filters: args,
	})
	if err != nil {
		return nil, err
	}
	if len(services) == 0 {
		return nil, nil
	}
	return &services[0], nil
}

func (d *daemon) createApplication(application core.Application) error {
	var replicas = uint64(1)
	_, err := d.dockerCli.ServiceCreate(context.Background(), swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: application.Name + "_" + application.ID,
			Labels: map[string]string{
				"zaap-app-id":   application.ID,
				"zaap-app-name": application.Name,
			},
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: application.Name,
				Env:   convertEnvironmentVariables(application.Environment),
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

func (d *daemon) updateApplication(application core.Application, service *swarm.Service) error {
	var replicas = uint64(1)
	_, err := d.dockerCli.ServiceUpdate(context.Background(), service.ID, service.Version, swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: application.Name + "_" + application.ID,
			Labels: map[string]string{
				"zaap-app-id":   application.ID,
				"zaap-app-name": application.Name,
			},
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: application.Name,
				Env:   convertEnvironmentVariables(application.Environment),
			},
		},
		Mode: swarm.ServiceMode{
			Replicated: &swarm.ReplicatedService{
				Replicas: &replicas,
			},
		},
	}, types.ServiceUpdateOptions{})
	return err
}

func (d *daemon) deployApplication(payload core.DeploymentPayload) error {
	log := logrus.WithField("application", payload.Application.ID)
	log.Info("deployment requested")
	currentApp, err := d.getApplication(payload.Application.ID)
	if err != nil {
		return err
	}

	if currentApp == nil {
		log.Info("application does not exists, creating")
		err = d.createApplication(payload.Application)
	} else {
		log.Info("application already exists, updating")
		err = d.updateApplication(payload.Application, currentApp)
	}
	return err
}

func convertEnvironmentVariables(environment map[string]string) []string {
	var env []string
	for key, value := range environment {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}

	return env
}
