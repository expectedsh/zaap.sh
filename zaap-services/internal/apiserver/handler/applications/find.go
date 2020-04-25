package applications

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"net/http"
)

func HandleFind(deploymentStore core.DeploymentStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		application := request.ApplicationFrom(r.Context())

		deployments, err := deploymentStore.ListByApplication(r.Context(), application.ID)
		if err != nil {
			response.InternalServerError(w)
			return
		}
		application.Deployments = *deployments

		response.Ok(w, map[string]interface{}{
			"application": application,
		})
	}
}
