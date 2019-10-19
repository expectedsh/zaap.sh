package httpx

import (
  "context"
  "encoding/json"
  "github.com/gorilla/mux"
  "net/http"
  "sync"
)

type Context struct {
  Request  *Request
  Response *Response
  logger   *ContextLogger
}

type ContextLogger struct {
  mutex  *sync.RWMutex
  fields map[string]interface{}
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
  return &Context{
    Request: &Request{
      Request: r,
    },
    Response: &Response{
      ResponseWriter: w,
    },
    logger: &ContextLogger{
      mutex:  &sync.RWMutex{},
      fields: map[string]interface{}{},
    },
  }
}

func (c *Context) Context() context.Context {
  return c.Request.Context()
}

func (c *Context) QueryParam(key string) string {
  return c.Request.FormValue(key)
}

func (c *Context) PathParam(key string) string {
  return mux.Vars(c.Request.Request)[key]
}

func (c *Context) Get(key string) interface{} {
  return c.Context().Value(key)
}

func (c *Context) LogFields() map[string]interface{} {
  c.logger.mutex.RLock()
  fields := c.logger.fields
  c.logger.mutex.RUnlock()
  return fields
}

func (c *Context) Set(key string, value interface{}) *Context {
  ctx := context.WithValue(c.Context(), key, value)
  c.Request.WithContext(ctx)
  return c
}

func (c *Context) Log(key string, v interface{}) *Context {
  c.logger.mutex.Lock()
  c.logger.fields[key] = v
  c.logger.mutex.Unlock()
  return c
}

func (c *Context) StatusCode(statusCode int) *Context {
  c.Response.WriteHeader(statusCode)
  return c
}

func (c *Context) Header(key, value string) *Context {
  c.Response.Header().Set(key, value)
  return c
}

func (c *Context) Write(b []byte) {
  _, _ = c.Response.Write(b)
}

func (c *Context) Redirect(url string, code int) {
  http.Redirect(c.Response, c.Request.Request, url, code)
}

func (c *Context) Json(statusCode int, v interface{}) {
  b, _ := json.Marshal(v)
  c.
    StatusCode(statusCode).
    Header("Content-Type", "application/json; charset=utf-8").
    Write(b)
}

func (c *Context) Resource(name string, v interface{}) {
  c.Json(http.StatusOK, map[string]interface{}{
    name: v,
  })
}

func (c *Context) Error(statusCode int, message string) {
  c.Json(statusCode, map[string]interface{}{
    "error": map[string]interface{}{
      "message": message,
    },
  })
}

func (c *Context) ErrorInternal(err error) {
  c.Log("error", err)
  c.Error(http.StatusInternalServerError, "Something went wrong.")
}

func (c *Context) ErrorNotFound() {
  c.Error(http.StatusNotFound, "Resource not found.")
}

func (c *Context) ErrorForbidden(w http.ResponseWriter) {
  c.Error(http.StatusForbidden, "You do not have access for the attempted action.")
}

func (c *Context) ErrorBadRequest(message string, fields map[string]string) {
  c.Json(http.StatusBadRequest, map[string]interface{}{
    "error": map[string]interface{}{
      "message": message,
      "fields":  fields,
    },
  })
}
