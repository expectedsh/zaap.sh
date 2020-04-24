package notifiers

import "github.com/expected.sh/zaap.sh/zaap-services/pkg/core"

type Notifier interface {
	WhenApplicationDeploymentRequest(application *core.Application) error

	WhenApplicationDeleted(id, name string) error
}
