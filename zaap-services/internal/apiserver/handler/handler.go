package handler

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/config"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/handler/users"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/service"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/store"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

func New(config *config.Config, db *gorm.DB) chi.Router {
	r := chi.NewRouter()

	userStore := store.NewUserStore(db)
	userService := service.NewUserService(config.SecretKey)

	r.Post("/users", users.HandleCreate(userStore, userService))

	return r
}
