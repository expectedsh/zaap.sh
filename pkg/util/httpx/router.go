package httpx

import (
  "github.com/gorilla/mux"
  "github.com/sirupsen/logrus"
  "net/http"
  "time"
)

type Router struct {
  http.Handler
  router *mux.Router
}

type HandlerFunc func(ctx *Context)

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  r.router.ServeHTTP(w, req)
}

func (r Router) Get(path string, handler HandlerFunc) {
  r.router.HandleFunc(path, wrapHandler(handler)).Methods("GET")
}

func (r Router) Post(path string, handler HandlerFunc) {
  r.router.HandleFunc(path, wrapHandler(handler)).Methods("POST")
}

func (r Router) Put(path string, handler HandlerFunc) {
  r.router.HandleFunc(path, wrapHandler(handler)).Methods("PUT")
}

func (r Router) Patch(path string, handler HandlerFunc) {
  r.router.HandleFunc(path, wrapHandler(handler)).Methods("PATCH")
}

func (r Router) Delete(path string, handler HandlerFunc) {
  r.router.HandleFunc(path, wrapHandler(handler)).Methods("DELETE")
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
