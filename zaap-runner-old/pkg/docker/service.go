package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"io"
)

func (c *Client) ServiceGetFromApplication(ctx context.Context, id string) (*swarm.Service, error) {
	args := filters.NewArgs()
	args.Add("label", fmt.Sprintf("zaap-app-id=%s", id))
	services, err := c.client.ServiceList(ctx, types.ServiceListOptions{
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

func (c *Client) ServiceCreate(ctx context.Context, spec swarm.ServiceSpec) error {
	_, err := c.client.ServiceCreate(ctx, spec, types.ServiceCreateOptions{})
	return err
}

func (c *Client) ServiceUpdate(ctx context.Context, spec swarm.ServiceSpec, service *swarm.Service) error {
	_, err := c.client.ServiceUpdate(ctx, service.ID, service.Version, spec, types.ServiceUpdateOptions{})
	return err
}

func (c *Client) ServiceDelete(ctx context.Context, service *swarm.Service) error {
	return c.client.ServiceRemove(ctx, service.ID)
}

func (c *Client) ServiceGetLogs(ctx context.Context, service *swarm.Service) (io.ReadCloser, error) {
	return c.client.ServiceLogs(ctx, service.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Details:    true,
		Timestamps: true,
	})
}
