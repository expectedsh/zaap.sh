package app

import (
  "context"
  "github.com/docker/docker/api/types"
  "github.com/docker/docker/api/types/swarm"
  "github.com/remicaumette/zaap.sh/pkg/protocol"
)

func (s *Service) CreateApp(ctx context.Context, r *protocol.CreateAppRequest) (*protocol.App, error) {
  replicas := uint64(1)
  resources := &swarm.Resources{}
  if r.Memory != -1 {
    resources.MemoryBytes = int64(r.Memory * 1024 * 1024)
  }
  if r.Cpu != -1 {
    resources.NanoCPUs = int64(r.Cpu * 100000000)
  }
  _, err := s.Docker.ServiceCreate(ctx, swarm.ServiceSpec{
    Annotations: swarm.Annotations{
      Name: r.Name,
    },
    TaskTemplate: swarm.TaskSpec{
      ContainerSpec: swarm.ContainerSpec{
        Image: r.Image,
        Env:   r.Env,
      },
      Resources: &swarm.ResourceRequirements{
        Limits:       resources,
        Reservations: resources,
      },
    },
    Mode: swarm.ServiceMode{
      Replicated: &swarm.ReplicatedService{
        Replicas: &replicas,
      },
    },
  }, types.ServiceCreateOptions{})
  if err != nil {
    return nil, err
  }
  return &protocol.App{
    Name: r.Name,
  }, nil
}
