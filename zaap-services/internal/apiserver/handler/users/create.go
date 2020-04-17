package users

import (
	"encoding/json"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/sirupsen/logrus"
	"net/http"
)

type (
	createUserRequest struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
	}

	createUserResponse struct {
		Token string     `json:"token"`
		User  *core.User `json:"user"`
	}
)

func HandleCreate(store core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := new(createUserRequest)
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			logrus.WithError(err).Error("could not decode request payload")
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		user := &core.User{
			Email:     in.Email,
			Password:  in.Password,
			FirstName: in.FirstName,
		}
		if err := store.Create(r.Context(), user); err != nil {
			logrus.WithError(err).Error("could not create user")
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(&createUserResponse{
			Token: "",
			User: user,
		})
	}
}
