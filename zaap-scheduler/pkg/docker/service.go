package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
)

func (d *Docker) ServiceGetFromApplication(ctx context.Context, id string) (*swarm.Service, error) {
	args := filters.NewArgs()
	args.Add("label", fmt.Sprintf("zaap-app-id=%s", id))
	services, err := d.Client.ServiceList(ctx, types.ServiceListOptions{
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

func (d *Docker) ServiceCreate(ctx context.Context, spec swarm.ServiceSpec) error {
	_, err := d.Client.ServiceCreate(ctx, spec, types.ServiceCreateOptions{})
	return err
}

func (d *Docker) ServiceUpdate(ctx context.Context, spec swarm.ServiceSpec, service *swarm.Service) error {
	_, err := d.Client.ServiceUpdate(ctx, service.ID, service.Version, spec, types.ServiceUpdateOptions{})
	return err
}

func (d *Docker) ServiceDelete(ctx context.Context, service *swarm.Service) error {
	return d.Client.ServiceRemove(ctx, service.ID)
}
