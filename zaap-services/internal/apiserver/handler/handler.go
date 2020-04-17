package handler

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/handler/users"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/store"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

func New(db *gorm.DB) chi.Router {
	r := chi.NewRouter()

	userStore := store.NewUserStore(db)

	r.Post("/users", users.HandleCreate(userStore))

	return r
}
