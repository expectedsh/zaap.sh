package runner

import (
	"context"
	"errors"
	"github.com/expected.sh/zaap.sh/zaap-runner/internal/runner/config"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/kubernetes"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

var ErrUnauthorized = errors.New("unauthorized")

type Runner struct {
	config     *config.Config
	client     *kubernetes.Client
	grpcServer *grpc.Server
}

func New(config *config.Config) *Runner {
	return &Runner{config: config}
}

func (r *Runner) Start() error {
	logrus.Info("starting runner")

	kClient, err := r.config.KubernetesClient()
	if err != nil {
		return err
	}
	r.client = kClient

	r.grpcServer = grpc.NewServer(grpc.ChainUnaryInterceptor(r.AuthHandler))
	runnerpb.RegisterRunnerServer(r.grpcServer, r)

	lis, err := net.Listen("tcp", r.config.Addr)
	if err != nil {
		return err
	}

	logrus.Infof("listening on %v", r.config.Addr)
	return r.grpcServer.Serve(lis)
}

func (r *Runner) Shutdown() {
	r.grpcServer.GracefulStop()
}

func (r *Runner) AuthHandler(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrUnauthorized
	}

	header, ok := md["authorization"]
	if !ok || len(header) == 0 || header[0] != r.config.Token {
		return nil, ErrUnauthorized
	}

	return handler(ctx, req)
}
