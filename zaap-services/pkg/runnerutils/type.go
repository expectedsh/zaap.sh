package runnerutils

import (
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
)

func ConvertType(runnerType protocol.RunnerType) core.RunnerType {
	switch runnerType {
	case protocol.RunnerType_DOCKER_SWARM:
		return core.RunnerTypeDockerSwarm
	case protocol.RunnerType_KUBERNETES:
		return core.RunnerTypeKubernetes
	default:
		return core.RunnerTypeUnknown
	}
}
