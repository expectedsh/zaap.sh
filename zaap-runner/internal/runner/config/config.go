package config

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/kubernetes"
	"github.com/kelseyhightower/envconfig"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Config struct {
	Addr          string   `required:"true" envconfig:"ADDR" default:":8090"`
	Token         string   `required:"true" envconfig:"TOKEN"`
	KubeConfig    string   `envconfig:"KUBECONFIG"`
	KubeNamespace string   `required:"true" envconfig:"NAMESPACE" default:"zaap"`
	ExternalIps   []string `required:"true" split_words:"true" envconfig:"EXTERNAL_IPS"`
}

func FromEnv() (*Config, error) {
	config := new(Config)
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}

func (c Config) KubernetesClient() (*kubernetes.Client, error) {
	var kConfig *rest.Config
	if c.KubeConfig != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", c.KubeConfig)
		if err != nil {
			return nil, err
		}
		kConfig = cfg
	} else {
		cfg, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
		kConfig = cfg
	}
	return kubernetes.NewClient(c.KubeNamespace, kConfig)
}
