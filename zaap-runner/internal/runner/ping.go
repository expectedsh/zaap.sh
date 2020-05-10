package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
)

func (r *Runner) Ping(_ context.Context, req *runnerpb.PingRequest) (*runnerpb.PingResponse, error) {
	return &runnerpb.PingResponse{
		Time: req.Time,
	}, nil
}
