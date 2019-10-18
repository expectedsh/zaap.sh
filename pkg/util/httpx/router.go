package httpx

import (
  "github.com/gorilla/mux"
  "github.com/sirupsen/logrus"
  "net/http"
  "time"
)

type Router struct {
  http.Handler
  Router *mux.Router
}

type HandlerFunc func(ctx *Context)

type NextFunc func()

type MiddlewareFunc func(ctx *Context, next NextFunc)

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  r.Router.ServeHTTP(w, req)
}

func (r *Router) Get(path string, handler HandlerFunc) {
  r.Router.HandleFunc(path, wrapHandler(handler)).Methods("GET")
}

func (r *Router) Post(path string, handler HandlerFunc) {
  r.Router.HandleFunc(path, wrapHandler(handler)).Methods("POST")
}

func (r *Router) Put(path string, handler HandlerFunc) {
  r.Router.HandleFunc(path, wrapHandler(handler)).Methods("PUT")
}

func (r *Router) Patch(path string, handler HandlerFunc) {
  r.Router.HandleFunc(path, wrapHandler(handler)).Methods("PATCH")
}

func (r *Router) Delete(path string, handler HandlerFunc) {
  r.Router.HandleFunc(path, wrapHandler(handler)).Methods("DELETE")
}

func (r *Router) Use(middleware MiddlewareFunc) {
  r.Router.Use(func(next http.Handler) http.Handler {
    return wrapHandler(func(ctx *Context) {
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

func wrapHandler(handler HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    startTime := time.Now()
    ctx := NewContext(w, r)
    handler(ctx)
    logrus.
      WithField("duration", time.Now().Sub(startTime)).
      WithField("http_method", r.Method).
      WithField("http_path", r.RequestURI).
      WithField("http_status", ctx.Response.StatusCode()).
      WithField("remote_addr", ctx.Request.RemoteAddr).
      Info("canonical-log-line")
  }
}
