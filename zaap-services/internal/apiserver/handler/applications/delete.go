package applications

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func HandleDelete(store core.ApplicationStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.FromString(chi.URLParam(r, "id"))
		if err != nil {
			response.NotFound(w)
			return
		}

		if err := store.Delete(r.Context(), id); err != nil {
			response.InternalServerError(w)
			return
		}

		response.NoContent(w)
	}
}
