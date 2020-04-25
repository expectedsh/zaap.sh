package runners

import (
	"encoding/json"
	"errors"
	"github.com/expected.sh/zaap.sh/zaap-scheduler/pkg/protocol"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type createRunnerRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Token       string `json:"token"`
}

func HandleCreate(store core.RunnerStore, service core.RunnerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := request.UserFrom(r.Context())

		in := new(createRunnerRequest)
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			response.BadRequest(w)
			return
		}

		runner := &core.Runner{
			Name:   in.Name,
			Url:    in.Url,
			Token:  in.Token,
			Status: core.RunnerStatusUnknown,
			Type:   core.RunnerTypeDockerSwarm,
			User:   user,
		}
		if in.Description != "" {
			runner.Description = &in.Description
		}

		if err := runner.Validate(); err != nil {
			response.UnprocessableEntity(w, err)
			return
		}

		client, conn, err := service.NewConnection(runner)
		if err != nil {
			response.UnprocessableEntity(w, validation.Errors{
				"url": err,
			})
			return
		}
		defer conn.Close()

		res, err := client.TestConnection(r.Context(), &protocol.TestConnectionRequest{
			Token: runner.Token,
		})
		if err != nil {
			response.UnprocessableEntity(w, validation.Errors{
				"url": err,
			})
			return
		} else if !res.Ok {
			response.UnprocessableEntity(w, validation.Errors{
				"token": errors.New("bad token"),
			})
			return
		}

		if err := store.Create(r.Context(), runner); err != nil {
			logrus.WithError(err).Warn("could not create runner")
			response.InternalServerError(w)
			return
		}

		response.Created(w, map[string]interface{}{
			"runner": runner,
		})
	}
}
