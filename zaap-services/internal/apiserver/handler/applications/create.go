package applications

import (
	"encoding/json"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type createApplicationRequest struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (r *createApplicationRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required, validation.Length(1, 0)),
		validation.Field(&r.Image, validation.Required, validation.Length(1, 0)),
	)
}

func HandleCreate(store core.ApplicationStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := request.UserFrom(r.Context())

		in := new(createApplicationRequest)
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			response.UnprocessableEntity(w, err)
			return
		}

		if err := in.Validate(); err != nil {
			response.UnprocessableEntity(w, err)
			return
		}

		deployment := &core.Deployment{
			ID:       uuid.NewV4(),
			Image:    in.Image,
			Replicas: 1,
		}
		application := &core.Application{
			Name:                in.Name,
			State:               core.ApplicationStateUnknown,
			UserID:              user.ID,
			CurrentDeploymentID: deployment.ID,
			Deployments:         []*core.Deployment{deployment},
		}
		if err := store.Create(r.Context(), application); err != nil {
			logrus.WithError(err).Warn("could not create application")
			response.InternalServerError(w)
			return
		}
		// @todo: deploy

		response.Created(w, map[string]interface{}{
			"application": application,
		})
	}
}
