package service

import (
	"github.com/expected.sh/zaap.sh/zaap-scheduler/pkg/protocol"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type runnerService struct {
	amqpConn  *amqp.Connection
}

func NewRunnerService(amqpConn *amqp.Connection) core.RunnerService {
	return &runnerService{
		amqpConn:  amqpConn,
	}
}

func (s runnerService) NewConnection(runner *core.Runner) (protocol.SchedulerClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(runner.Url, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	return protocol.NewSchedulerClient(conn), conn, nil
}
