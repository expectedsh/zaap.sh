package handler

import (
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/config"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/handler/applications"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/handler/auth"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/handler/me"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver/handler/users"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/service"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jinzhu/gorm"
)

func New(config *config.Config, db *gorm.DB) chi.Router {
	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler)

	userStore := store.NewUserStore(db)
	userService := service.NewUserService(config.SecretKey)

	applicationStore := store.NewApplicationStore(db)

	r.Post("/auth/login", auth.HandleLogin(userStore, userService))
	r.Post("/users", users.HandleCreate(userStore, userService))

	r.Group(func(r chi.Router) {
		r.Use(auth.Required(userStore, userService))

		r.Get("/me", me.HandleFind())
		r.Patch("/me", me.HandleUpdate(userStore))

		r.Get("/applications", applications.HandleList(applicationStore))
		r.Post("/applications", applications.HandleCreate(applicationStore))
		r.Get("/applications/{id}", applications.HandleFind(applicationStore))
		r.Delete("/applications/{id}", applications.HandleDelete(applicationStore))
	})

	return r
}
