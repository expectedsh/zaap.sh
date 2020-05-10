package request

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
)

const RunnerKey = "runner"

func WithRunner(ctx context.Context, runner *core.Runner) context.Context {
	return context.WithValue(ctx, RunnerKey, runner)
}

func RunnerFrom(ctx context.Context) *core.Runner {
	runner, _ := ctx.Value(RunnerKey).(*core.Runner)
	return runner
}
