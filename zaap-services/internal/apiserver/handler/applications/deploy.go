package applications

import (
	"github.com/expected.sh/zaap.sh/zaap-scheduler/pkg/protocol"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"net/http"
)

func HandleDeploy(deploymentStore core.DeploymentStore, userService core.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		application := request.ApplicationFrom(r.Context())

		client, conn, err := userService.NewSchedulerConnection(application.User)
		if err != nil {
			response.InternalServerError(w)
			return
		}
		defer conn.Close()

		currentDeployment, err := deploymentStore.Find(r.Context(), application.CurrentDeploymentID)
		if err != nil {
			response.InternalServerError(w)
			return
		} else if currentDeployment == nil {
			response.NotFound(w)
			return
		}

		_, err = client.DeployApplication(r.Context(), &protocol.DeployApplicationRequest{
			Application: &protocol.Application{
				Id:          application.ID.String(),
				Name:        application.Name,
				Image:       currentDeployment.Image,
				Replicas:    uint32(currentDeployment.Replicas),
				Environment: currentDeployment.Environment,
			},
		})
		if err != nil {
			response.InternalServerError(w)
			return
		}

		response.Ok(w, map[string]interface{}{
			"application": application,
		})
	}
}
