package request

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
)

const ApplicationKey = "application"

func WithApplication(ctx context.Context, application *core.Application) context.Context {
	return context.WithValue(ctx, ApplicationKey, application)
}

func ApplicationFrom(ctx context.Context) *core.Application {
	application, _ := ctx.Value(ApplicationKey).(*core.Application)
	return application
}
