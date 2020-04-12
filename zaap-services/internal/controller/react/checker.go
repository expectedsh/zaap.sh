package react

import (
	"github.com/docker/docker/api/types/events"
)

func IsReactiveTo(event events.Message) bool {
	for key, value := range reactionDefinition {
		if event.Type == key {
			if _, ok := value[event.Action]; ok {
				return true
			}
		}
	}

	return false
}
