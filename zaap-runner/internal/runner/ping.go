package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
)

func (r *Runner) Ping(_ context.Context, req *protocol.PingRequest) (*protocol.PingResponse, error) {
	return &protocol.PingResponse{
		Time: req.Time,
	}, nil
}
