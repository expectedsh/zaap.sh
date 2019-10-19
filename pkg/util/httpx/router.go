package httpx

import (
  "github.com/gorilla/mux"
  "net/http"
)

type Router struct {
  http.Handler
  Router *mux.Router
}

type HandlerFunc interface{}

type NextFunc func()

type MiddlewareFunc func(ctx *Context, next NextFunc)

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  r.Router.ServeHTTP(w, req)
}

func (r *Router) Get(path string, handler HandlerFunc) {
  r.Router.HandleFunc(path, WrapHandler(handler)).Methods("GET")
}

func (r *Router) Post(path string, handler HandlerFunc) {
  r.Router.HandleFunc(path, WrapHandler(handler)).Methods("POST")
}

func (r *Router) Put(path string, handler HandlerFunc) {
  r.Router.HandleFunc(path, WrapHandler(handler)).Methods("PUT")
}

func (r *Router) Patch(path string, handler HandlerFunc) {
  r.Router.HandleFunc(path, WrapHandler(handler)).Methods("PATCH")
}

func (r *Router) Delete(path string, handler HandlerFunc) {
  r.Router.HandleFunc(path, WrapHandler(handler)).Methods("DELETE")
}

func (r *Router) Use(middleware MiddlewareFunc) {
  r.Router.Use(func(next http.Handler) http.Handler {
    return WrapHandler(func(ctx *Context) {
      middleware(ctx, func() {
        next.ServeHTTP(ctx.Response, ctx.Request.Request)
      })
    })
  })
}

func (r *Router) Group(path string) *Router {
  return &Router{
    Router: r.Router.PathPrefix(path).Subrouter(),
  }
}
