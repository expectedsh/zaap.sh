package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
)

func (r *Runner) Ping(_ context.Context, req *runnerpb.PingRequest) (*runnerpb.PingReply, error) {
	return &runnerpb.PingReply{
		Time: req.Time,
	}, nil
}
