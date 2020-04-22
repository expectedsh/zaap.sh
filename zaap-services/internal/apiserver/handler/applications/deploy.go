package applications

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HandleDeploy(service core.ApplicationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		application := request.ApplicationFrom(r.Context())

		if err := service.Deploy(application); err != nil {
			logrus.WithError(err).Warn("could not deploy application")
			response.InternalServerError(w)
			return
		}

		logrus.Info(service.NotifyDeletion(application))

		response.Ok(w, map[string]interface{}{
			"application": application,
		})
	}
}
