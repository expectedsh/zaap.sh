package auth

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"net/http"
	"strings"
)

func Required(store core.UserStore, service core.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("authorization")
			headerParts := strings.Split(header, " ")
			token := headerParts[len(headerParts)-1]

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
