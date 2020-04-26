package applications

import (
	"encoding/json"
	"errors"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type createApplicationRequest struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	RunnerID string `json:"runner_id"`
}

func (r *createApplicationRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(
			&r.Name,
			validation.Required,
			validation.Length(3, 50),
			validation.
				Match(core.ApplicationNameRegex).
				Error("should only contain letters, numbers, and dashes"),
		),
		validation.Field(
			&r.Image,
			validation.Required,
			validation.
				Match(core.DeploymentImageRegex).
				Error("invalid image"),
		),
		validation.Field(&r.RunnerID, validation.Required, is.UUIDv4),
	)
}

func HandleCreate(store core.ApplicationStore, runnerStore core.RunnerStore, service core.ApplicationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := request.UserFrom(r.Context())

		in := new(createApplicationRequest)
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			response.BadRequest(w)
			return
		} else if err := in.Validate(); err != nil {
			response.UnprocessableEntity(w, err)
			return
		}

		runner, err := runnerStore.Find(r.Context(), uuid.FromStringOrNil(in.RunnerID))
		if err != nil {
			response.InternalServerError(w)
			return
		} else if runner == nil || runner.UserID != user.ID {
			response.UnprocessableEntity(w, validation.Errors{
				"runner_id": errors.New("not found"),
			})
			return
		}

		deployment := &core.Deployment{
			ID:       uuid.NewV4(),
			Image:    in.Image,
			Replicas: 1,
		}
		application := &core.Application{
			Name:                in.Name,
			Status:              core.ApplicationStatusUnknown,
			CurrentDeploymentID: deployment.ID,
			User:                user,
			Runner:              runner,
			Deployments:         []*core.Deployment{deployment},
		}

		if err := store.Create(r.Context(), application); err != nil {
			logrus.WithError(err).Warn("could not create application")
			response.InternalServerError(w)
			return
		}

		if err := service.NotifyCreated(application); err != nil {
			logrus.WithError(err).Warn("could not notify created")
		}

		response.Created(w, map[string]interface{}{
			"application": application,
		})
	}
}
