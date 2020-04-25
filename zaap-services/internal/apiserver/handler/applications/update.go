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

type updateApplicationRequest struct {
	Name        *string           `json:"name"`
	Image       *string           `json:"image"`
	Replicas    *int              `json:"replicas"`
	Environment *core.Environment `json:"environment"`
}

func (r updateApplicationRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Length(1, 0)),
		validation.Field(&r.Image, validation.Length(1, 0)),
		validation.Field(&r.Replicas, validation.Min(1)),
	)
}

func HandleUpdate(store core.ApplicationStore, deploymentStore core.DeploymentStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		application := request.ApplicationFrom(r.Context())

		currentDeployment, err := deploymentStore.Find(r.Context(), application.CurrentDeploymentID)
		if err != nil {
			response.InternalServerError(w)
			return
		} else if currentDeployment == nil {
			response.NotFound(w)
			return
		}

		in := new(updateApplicationRequest)
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			response.BadRequest(w)
			return
		} else if err := in.Validate(); err != nil {
			response.UnprocessableEntity(w, err)
			return
		}

		if in.Name != nil {
			application.Name = *in.Name
		}
		if in.Image != nil || in.Replicas != nil || in.Environment != nil {
			deployment := &core.Deployment{
				ID:          uuid.NewV4(),
				Image:       currentDeployment.Image,
				Replicas:    currentDeployment.Replicas,
				Environment: currentDeployment.Environment,
			}

			if in.Image != nil {
				deployment.Image = *in.Image
			}
			if in.Replicas != nil {
				deployment.Replicas = *in.Replicas
			}
			if in.Environment != nil {
				deployment.Environment = *in.Environment
			}

			application.Deployments = append(application.Deployments, deployment)
			application.CurrentDeploymentID = deployment.ID
		}

		if err := store.Update(r.Context(), application); err != nil {
			logrus.WithError(err).Warn("could not update application")
			response.InternalServerError(w)
			return
		}

		response.Ok(w, map[string]interface{}{
			"application": application,
		})
	}
}
