package me

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/request"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/response"
	"net/http"
)

func HandleFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Ok(w, map[string]interface{}{
			"user": request.UserFrom(r.Context()),
		})
	}
}
