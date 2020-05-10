package applications

import (
	"encoding/json"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type updateApplicationRequest struct {
	Image            *string           `json:"image"`
	Replicas         *int              `json:"replicas"`
	Environment      *core.Environment `json:"environment"`
	Domains          []string          `json:"domains"`
	Roles            []string          `json:"roles"`
	ImagePullSecrets []string          `json:"image_pull_secrets"`
}

func (r *updateApplicationRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(
			&r.Image,
			validation.
				Match(core.DeploymentImageRegex).
				Error("invalid image"),
		),
		validation.Field(&r.Replicas, validation.Min(1)),
		validation.Field(&r.Domains, validation.Each(is.Domain)),
	)
}

func HandleUpdate(store core.ApplicationStore, deploymentStore core.DeploymentStore, service core.ApplicationService) http.HandlerFunc {
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
			logrus.Info(err)
			response.UnprocessableEntity(w, err)
			return
		}

		if in.Image != nil || in.Replicas != nil || in.Environment != nil || in.Roles != nil || in.ImagePullSecrets != nil {
			deployment := &core.Deployment{
				ID:               uuid.NewV4(),
				Image:            currentDeployment.Image,
				Replicas:         currentDeployment.Replicas,
				Environment:      currentDeployment.Environment,
				Roles:            currentDeployment.Roles,
				ImagePullSecrets: currentDeployment.ImagePullSecrets,
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
			if in.Roles != nil {
				deployment.Roles = removeDuplicates(in.Roles)
			}
			if in.ImagePullSecrets != nil {
				deployment.ImagePullSecrets = removeDuplicates(in.ImagePullSecrets)
			}

			application.Deployments = append(application.Deployments, deployment)
			application.CurrentDeploymentID = deployment.ID
		}
		if in.Domains != nil {
			application.Domains = removeDuplicates(in.Domains)
		}

		if err := store.Update(r.Context(), application); err != nil {
			logrus.WithError(err).Warn("could not update application")
			response.InternalServerError(w)
			return
		}

		if err := service.NotifyUpdated(application); err != nil {
			logrus.WithError(err).Warn("could not notify updated")
		}

		response.Ok(w, map[string]interface{}{
			"application": application,
		})
	}
}

func removeDuplicates(slice []string) []string {
	var list []string
	keys := make(map[string]bool)
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
