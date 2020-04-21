package deployment_manager

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/deployment-manager/config"
)

type Manager struct {
	config *config.Config
}

func New(config *config.Config) *Manager {
	return &Manager{
		config: config,
	}
}

func (m *Manager) Start() error {
	return nil
}

func (m *Manager) Shutdown(ctx context.Context) error {
	return nil
}
