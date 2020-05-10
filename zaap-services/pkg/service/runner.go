package service

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/connector/rabbitmq"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/messaging"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var RunnerEventsExchange = messaging.NewDurableExchangeTopic("runner_events")

type runnerService struct {
	amqpConn  *rabbitmq.Connection
	publisher *messaging.Publisher
}

func NewRunnerService(amqpConn *rabbitmq.Connection) core.RunnerService {
	return &runnerService{
		amqpConn:  amqpConn,
		publisher: messaging.NewPublisher(amqpConn, RunnerEventsExchange),
	}
}

func withRunnerCredentials(runner *core.Runner) grpc.DialOption {
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(
			metadata.AppendToOutgoingContext(ctx, "authorization", runner.Token),
			method,
			req,
			reply,
			cc,
			opts...,
		)
	})
}

func (s runnerService) NewConnection(runner *core.Runner) (runnerpb.RunnerClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(runner.Url, grpc.WithInsecure(), withRunnerCredentials(runner))
	if err != nil {
		return nil, nil, err
	}
	return runnerpb.NewRunnerClient(conn), conn, nil
}

func (s runnerService) NotifyStatusChanged(runner *core.Runner) error {
	return s.publisher.Publish(messaging.DeliveryModeTransient, &protocol.RunnerStatusChanged{
		Id:     runner.ID.String(),
		Status: runner.Status.ToMessagingFormat(),
	})
}
