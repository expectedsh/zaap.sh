package notifiers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type DiscordNotifier struct {
	webhookUrl string
}

func NewDiscordNotifier(webhookUrl string) Notifier {
	return &DiscordNotifier{webhookUrl}
}

func (d DiscordNotifier) WhenApplicationDeleted(id, name string) error {
	b, err := json.Marshal(map[string]interface{}{
		"username": "zaap.sh",
		"content": fmt.Sprintf("Application %v (id: %v) deleted", name, id),
	})
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Post(d.webhookUrl, "application/json", bytes.NewBuffer(b))
	v, _ := ioutil.ReadAll(res.Body)
	logrus.Info(string(v))
	return err
}
