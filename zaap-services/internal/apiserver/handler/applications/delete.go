package applications

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HandleDelete(store core.ApplicationStore, service core.ApplicationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		application := request.ApplicationFrom(r.Context())

		if err := store.Delete(r.Context(), application.ID); err != nil {
			response.InternalServerError(w)
			return
		}

		if err := service.NotifyDeletion(application); err != nil {
			logrus.WithError(err).Warn("could not notify deletion")
			return
		}

		response.NoContent(w)
	}
}
