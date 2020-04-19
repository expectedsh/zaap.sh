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
	deploymentStore := store.NewDeploymentStore(db)

	r.Post("/auth/login", auth.HandleLogin(userStore, userService))
	r.Post("/users", users.HandleCreate(userStore, userService))

	r.Group(func(r chi.Router) {
		r.Use(AuthRequired(userStore, userService))

		r.Route("/me", func(r chi.Router) {
			r.Get("/", me.HandleFind())
			r.Patch("/", me.HandleUpdate(userStore, userService))
		})

		r.Route("/applications", func(r chi.Router) {
			r.Get("/", applications.HandleList(applicationStore))
			r.Post("/", applications.HandleCreate(applicationStore))

			r.Route("/{id}", func(r chi.Router) {
				r.Use(InjectApplication(applicationStore))

				r.Get("/", applications.HandleFind(deploymentStore))
				r.Patch("/", applications.HandleUpdate(applicationStore, deploymentStore))
				r.Delete("/", applications.HandleDelete(applicationStore))
				r.Get("/logs", applications.HandleLogs(userService))
				r.Post("/deploy", applications.HandleDeploy(deploymentStore, userService))
			})
		})
	})

	return r
}
