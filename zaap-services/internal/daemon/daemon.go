package daemon

import (
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"

	application "github.com/remicaumette/zaap.sh/zaap-services/pkg/models/applications"
)

type daemon struct {
	dockerCli *client.Client
}

func newDaemon(client *client.Client) *daemon {
	return &daemon{
		dockerCli: client,
	}
}

func (d *daemon) deployApplication(payload application.DeploymentPayload) error {
	logrus.WithField("daemon-func", "deployApplication").WithField("payload", payload).Info()

	return nil
}
