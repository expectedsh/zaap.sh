package service

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type runnerService struct {
	amqpConn *amqp.Connection
}

func NewRunnerService(amqpConn *amqp.Connection) core.RunnerService {
	return &runnerService{
		amqpConn: amqpConn,
	}
}

func withRunnerCredentials(runner *core.Runner) grpc.DialOption {
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		metadata.AppendToOutgoingContext(ctx, "authorization", runner.Token)
		return nil
	})
}

func (s runnerService) NewConnection(runner *core.Runner) (protocol.RunnerClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(runner.Url, grpc.WithInsecure(), withRunnerCredentials(runner))
	if err != nil {
		return nil, nil, err
	}
	return protocol.NewRunnerClient(conn), conn, nil
}
