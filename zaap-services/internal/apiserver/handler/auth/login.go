package auth

import (
	"encoding/json"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/sirupsen/logrus"
	"net/http"
)

type (
	loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	loginResponse struct {
		Token string `json:"token"`
	}
)

func HandleLogin(store core.UserStore, service core.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := new(loginRequest)
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			response.BadRequest(w)
			return
		}

		user, err := store.FindByEmail(r.Context(), in.Email)
		if err != nil {
			response.InternalServerError(w)
			return
		}

		if user == nil || !service.ComparePassword(user.Password, in.Password) {
			response.WriteResponse(w, http.StatusNotFound, map[string]string{
				"message": "Invalid email or password.",
			})
			return
		}

		token, err := service.IssueToken(user)
		if err != nil {
			logrus.WithError(err).Error("could not issue token")
			response.InternalServerError(w)
			return
		}

		response.Created(w, &loginResponse{Token: token})
	}
}
