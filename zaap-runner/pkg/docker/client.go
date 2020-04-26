package docker

import "github.com/docker/docker/client"

type Client struct {
	client *client.Client
}

func NewClient(client *client.Client) *Client {
	return &Client{client: client}
}
