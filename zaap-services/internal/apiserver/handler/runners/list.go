package runners

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"net/http"
)

func HandleList(store core.RunnerStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := request.UserFrom(r.Context())

		runners, err := store.ListByUser(r.Context(), user.ID)
		if err != nil {
			response.InternalServerError(w)
			return
		}

		response.Ok(w, map[string]interface{}{
			"runners": runners,
		})
	}
}
