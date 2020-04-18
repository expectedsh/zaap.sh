package applications

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"net/http"
)

func HandleUpdate(store core.ApplicationStore, deploymentStore core.DeploymentStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		application := request.ApplicationFrom(r.Context())

		response.Ok(w, map[string]interface{}{
			"application": application,
		})
	}
}
