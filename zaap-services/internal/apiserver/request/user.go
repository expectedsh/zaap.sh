package request

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
)

const UserKey = "user"

func WithUser(ctx context.Context, user *core.User) context.Context {
	return context.WithValue(ctx, UserKey, user)
}

func UserFrom(ctx context.Context) *core.User {
	user, _ := ctx.Value(UserKey).(*core.User)
	return user
}
