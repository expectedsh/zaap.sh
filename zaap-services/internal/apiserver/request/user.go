package request

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
)

const (
	UserKey = "user"
	ApplicationKey = "application"
)

func WithUser(ctx context.Context, user *core.User) context.Context {
	return context.WithValue(ctx, UserKey, user)
}

func UserFrom(ctx context.Context) *core.User {
	user, _ := ctx.Value(UserKey).(*core.User)
	return user
}

func WithApplication(ctx context.Context, application *core.Application) context.Context {
	return context.WithValue(ctx, ApplicationKey, application)
}

func ApplicationFrom(ctx context.Context) *core.Application {
	application, _ := ctx.Value(ApplicationKey).(*core.Application)
	return application
}
