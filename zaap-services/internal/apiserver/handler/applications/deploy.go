package applications

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"net/http"
)

func HandleDeploy(store core.ApplicationStore, userService core.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		application := request.ApplicationFrom(r.Context())

		//client, conn, err := userService.NewSchedulerConnection(application.User)
		//if err != nil {
		//	response.InternalServerError(w)
		//	return
		//}
		//defer conn.Close()
		//
		//_, err = client.DeployApplication(r.Context(), &protocol.DeployApplicationRequest{
		//	Application: &protocol.Application{
		//		Id:          application.ID.String(),
		//		Name:        application.Name,
		//		Image:       application.Image,
		//		Replicas:    uint32(application.Replicas),
		//		Environment: application.Environment,
		//	},
		//})
		//if err != nil {
		//	response.InternalServerError(w)
		//	return
		//}

		response.Ok(w, map[string]interface{}{
			"application": application,
		})
	}
}
