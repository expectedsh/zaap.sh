package notifiers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"net/http"
)

type DiscordNotifier struct {
	webhookUrl string
}

func NewDiscordNotifier(webhookUrl string) Notifier {
	return &DiscordNotifier{webhookUrl}
}

func (d DiscordNotifier) WhenApplicationDeploymentRequest(application *core.Application) error {
	b, err := json.Marshal(map[string]interface{}{
		"username": "zaap.sh",
		"content": fmt.Sprintf("Deployment requested for application `%v` (id: `%v`)", application.Name, application.ID),
	})
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Post(d.webhookUrl, "application/json", bytes.NewBuffer(b))
	return err
}

func (d DiscordNotifier) WhenApplicationDeleted(id, name string) error {
	b, err := json.Marshal(map[string]interface{}{
		"username": "zaap.sh",
		"content": fmt.Sprintf("Application `%v` (id: `%v`) deleted", name, id),
	})
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Post(d.webhookUrl, "application/json", bytes.NewBuffer(b))
	return err
}
