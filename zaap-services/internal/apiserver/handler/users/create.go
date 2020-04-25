package users

import (
	"encoding/json"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/sirupsen/logrus"
	"net/http"
)

type createUserRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
}

func (r *createUserRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 32)),
		validation.Field(&r.FirstName, validation.Required, validation.Length(1, 0)),
	)
}

func HandleCreate(store core.UserStore, service core.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := new(createUserRequest)
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			response.BadRequest(w)
			return
		} else if err := in.Validate(); err != nil {
			response.UnprocessableEntity(w, err)
			return
		}

		hashedPassword, err := service.HashPassword(in.Password)
		if err != nil {
			logrus.WithError(err).Error("could not hash password")
			response.InternalServerError(w)
			return
		}

		user := &core.User{
			Email:     in.Email,
			Password:  hashedPassword,
			FirstName: in.FirstName,
		}
		if err := store.Create(r.Context(), user); err != nil {
			logrus.WithError(err).Error("could not create user")
			response.InternalServerError(w)
			return
		}

		token, err := service.IssueToken(user)
		if err != nil {
			logrus.WithError(err).Error("could not issue token")
			response.InternalServerError(w)
			return
		}

		response.Created(w, map[string]interface{}{
			"token": token,
			"user":  user,
		})
	}
}
