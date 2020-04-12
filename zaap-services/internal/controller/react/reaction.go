package react

import "github.com/docker/docker/api/types/events"

type Reactor func(message events.Message)

var reactionDefinition = map[string]map[string]Reactor{
	"container": {
		"create":  nil,
		"destroy": nil,
		"die":     nil,
		"kill":    nil,
		"oom":     nil,
		"pause":   nil,
		"restart": nil,
		"start":   nil,
		"stop":    nil,
		"unpause": nil,
		"update":  nil,
	},

	"volumes": {
		"create":  nil,
		"destroy": nil,
		"mount":   nil,
		"unmount": nil,
	},

	"network": {
		"create":     nil,
		"connect":    nil,
		"destroy":    nil,
		"disconnect": nil,
		"remove":     nil,
	},
}

func On(message events.Message) {
	reactionDefinition[message.Type][message.Action](message)
}
