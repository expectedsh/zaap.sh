package httpx

import "github.com/gorilla/mux"

type Handler struct {
  Router
}

func NewHandler() *Handler {
  return &Handler{
    Router: Router{
      router: mux.NewRouter(),
    },
  }
}
