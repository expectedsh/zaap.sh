package me

import (
	"encoding/json"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"net/http"
)

type userUpdateRequest struct {
	Email        *string `json:"email"`
	FirstName    *string `json:"first_name"`
	SchedulerURL *string `json:"scheduler_url"`
}

func (r *userUpdateRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Email, is.Email),
		validation.Field(&r.FirstName, validation.Length(1, 0)),
		validation.Field(&r.SchedulerURL, is.URL),
	)
}

func HandleUpdate(store core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := request.UserFrom(r.Context())
		in := new(userUpdateRequest)

		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			response.UnprocessableEntity(w, err)
			return
		}

		if err := in.Validate(); err != nil {
			response.UnprocessableEntity(w, err)
			return
		}

		if in.Email != nil {
			user.Email = *in.Email
		}
		if in.FirstName != nil {
			user.FirstName = *in.FirstName
		}
		if in.SchedulerURL != nil {
			user.SchedulerURL = in.SchedulerURL
		}

		if err := store.Update(r.Context(), user); err != nil {
			response.InternalServerError(w)
			return
		}
		response.Ok(w, map[string]interface{}{
			"user": request.UserFrom(r.Context()),
		})
	}
}
