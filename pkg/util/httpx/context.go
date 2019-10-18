package httpx

import (
  "encoding/json"
  "net/http"
)

type Context struct {
  Request  *Request
  Response *Response
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
  return &Context{
    Request:  &Request{Request: r},
    Response: &Response{ResponseWriter: w},
  }
}

func (c *Context) Redirect(url string, code int) {
  http.Redirect(c.Response, c.Request.Request, url, code)
}

func (c *Context) Json(statusCode int, v interface{}) {
  b, _ := json.Marshal(v)
  c.Response.WriteHeader(statusCode)
  c.Response.Header().Set("Content-Type", "application/json; charset=utf-8")
  _, _ = c.Response.Write(b)
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
