package applications

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"net/http"
)

func HandleLogs(runnerStore core.RunnerStore, runnerService core.RunnerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		application := request.ApplicationFrom(r.Context())

		runner, err := runnerStore.Find(r.Context(), application.RunnerID)
		if err != nil {
			response.InternalServerError(w)
			return
		}

		client, conn, err := runnerService.NewConnection(runner)
		if err != nil {
			response.InternalServerError(w)
			return
		}
		defer conn.Close()

		logs, err := client.GetApplicationLogs(r.Context(), &runnerpb.GetApplicationLogsRequest{
			Id: application.ID.String(),
		})
		if err != nil {
			response.InternalServerError(w)
			return
		}
		defer logs.CloseSend()

		w.Header().Add("Content-Type", "text/event-stream")
		w.Header().Add("Connection", "keep-alive")

		for {
			log, err := logs.Recv()
			if err != nil {
				break
			}
			if err := sendLogLine(w, log); err != nil {
				break
			}
		}
	}
}

func sendLogLine(w http.ResponseWriter, log *runnerpb.GetApplicationLogsReply) error {
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
