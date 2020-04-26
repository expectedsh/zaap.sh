package runner

import (
	"github.com/docker/docker/client"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/docker"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Config struct {
	Addr        string   `required:"true" envconfig:"ADDR" default:":8090"`
	Token       string   `required:"true" envconfig:"EXTERNAL_IPS"`
	ExternalIps []string `required:"true" split_words:"true" envconfig:"EXTERNAL_IPS"`
}

type Runner struct {
	config       *Config
	dockerClient *docker.Client
	grpcServer   *grpc.Server
}

func ConfigFromEnv() (*Config, error) {
	config := new(Config)
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}

func New(config *Config) *Runner {
	return &Runner{config: config}
}

func (r *Runner) Start() error {
	logrus.Info("starting runner")

	dockerClient, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	defer dockerClient.Close()
	r.dockerClient = docker.NewClient(dockerClient)

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
