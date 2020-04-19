package applications

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/expected.sh/zaap.sh/zaap-scheduler/pkg/protocol"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HandleLogs(userService core.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		application := request.ApplicationFrom(r.Context())

		client, conn, err := userService.NewSchedulerConnection(application.User)
		if err != nil {
			response.InternalServerError(w)
			return
		}
		defer conn.Close()

		logs, err := client.GetApplicationLogs(r.Context(), &protocol.GetApplicationLogsRequest{Id: application.ID.String()})
		if err != nil {
			response.InternalServerError(w)
			return
		}
		defer logs.CloseSend()

		w.Header().Add("Content-Type", "text/event-stream")
		w.Header().Add("Connection", "keep-alive")

		for {
			log, err := logs.Recv()
			logrus.Info(log)
			if err != nil {
				logrus.WithError(err).Info("wtf")
				break
			}
			if err := sendLogLine(w, log); err != nil {
				logrus.WithError(err).Info("log line")
				break
			}
		}
	}
}

func sendLogLine(w http.ResponseWriter, log *protocol.GetApplicationLogsResponse) error {
	data, err := json.Marshal(log)
	if err != nil {
		return err
	}

	if _, err = fmt.Fprintf(w, "data: %s\n\n", data); err != nil {
		return err
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		return errors.New("flush unsupported")
	}
	flusher.Flush()

	return nil
}
