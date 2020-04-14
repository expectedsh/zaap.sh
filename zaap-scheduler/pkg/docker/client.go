package docker

import "github.com/docker/docker/client"

type Docker struct {
	Client *client.Client
}
