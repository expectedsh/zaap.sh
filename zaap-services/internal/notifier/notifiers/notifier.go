package notifier

import "github.com/expected.sh/zaap.sh/zaap-services/pkg/core"

type Notifier interface {
	WhenApplicationDeleted(application *core.Application) error
}
