package runners

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"net/http"
)

func HandleListClusterRoles(service core.RunnerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		runner := request.RunnerFrom(r.Context())

		client, conn, err := service.NewConnection(runner)
		if err != nil {
			response.InternalServerError(w)
			return
		}
		defer conn.Close()

		res, err := client.GetClusterRoles(r.Context(), &runnerpb.GetClusterRolesRequest{})
		if err != nil {
			response.InternalServerError(w)
			return
		}

		response.Ok(w, map[string]interface{}{
			"cluster_roles": res.Roles,
		})
	}
}
