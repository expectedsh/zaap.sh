package response

import (
	"encoding/json"
	"net/http"
)

func Ok(w http.ResponseWriter, v interface{}) {
	WriteResponse(w, http.StatusOK, v)
}

func Created(w http.ResponseWriter, v interface{}) {
	WriteResponse(w, http.StatusCreated, v)
}

func UnprocessableEntity(w http.ResponseWriter, err error) {
	WriteResponse(w, http.StatusUnprocessableEntity, err)
}

func InternalServerError(w http.ResponseWriter) {
	WriteResponse(w, http.StatusInternalServerError, map[string]string{
		"message": "Something went wrong.",
	})
}

func NotFound(w http.ResponseWriter) {
	WriteResponse(w, http.StatusNotFound, map[string]string{
		"message": "Resource not found.",
	})
}

func Forbidden(w http.ResponseWriter) {
	WriteResponse(w, http.StatusForbidden, map[string]string{
		"message": "You do not have access for the attempted action.",
	})
}

func NoContent(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
}

func WriteResponse(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
