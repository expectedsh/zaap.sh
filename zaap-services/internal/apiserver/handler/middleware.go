package handler

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strings"
)

func AuthRequired(store core.UserStore, service core.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := ""

			if query := r.URL.Query().Get("authorization"); query != "" {
				token = query
			} else {
				header := r.Header.Get("authorization")
				headerParts := strings.Split(header, " ")
				token = headerParts[len(headerParts)-1]
			}

			if token == "" {
				response.Forbidden(w)
				return
			}

			userId, err := service.UserIdFromToken(token)
			if err != nil {
				response.Forbidden(w)
				return
			}

			user, err := store.Find(r.Context(), *userId)
			if err != nil || user == nil {
				response.Forbidden(w)
				return
			}

			next.ServeHTTP(w, r.WithContext(
				request.WithUser(r.Context(), user),
			))
		})
	}
}

func InjectApplication(store core.ApplicationStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := request.UserFrom(r.Context())

			id, err := uuid.FromString(chi.URLParam(r, "id"))
			if err != nil {
				response.NotFound(w)
				return
			}

			application, err := store.Find(r.Context(), id)
			if err != nil {
				response.InternalServerError(w)
				return
			}

			if application == nil || user.ID.String() != application.UserID.String() {
				response.NotFound(w)
				return
			}

			next.ServeHTTP(w, r.WithContext(
				request.WithApplication(r.Context(), application),
			))
		})
	}
}
