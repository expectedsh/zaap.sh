package runner

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/internal/runner/config"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/kubernetes"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

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

	r.grpcServer = grpc.NewServer()
	protocol.RegisterRunnerServer(r.grpcServer, r)

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
