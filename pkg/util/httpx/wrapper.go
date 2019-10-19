package httpx

import (
  "encoding/json"
  "github.com/sirupsen/logrus"
  "io/ioutil"
  "net/http"
  "reflect"
  "time"
)

type Request struct {
  *http.Request
}

type Response struct {
  http.ResponseWriter
  statusCode int
}

func WrapHandler(handler HandlerFunc) http.HandlerFunc {
  f := reflect.ValueOf(handler)
  if f.Kind() != reflect.Func {
    panic("handler not a function")
  }
  t := f.Type()
  if t.NumIn() < 1 && t.NumIn() > 2 {
    panic("handler has invalid function parameter")
  }
  if !t.In(0).AssignableTo(reflect.TypeOf(&Context{})) {
    panic("first function parameter must be a *httpx.Context")
  }
  if t.NumIn() == 2 && t.In(1).Kind() != reflect.Struct {
    panic("second function parameter must be a structure")
  }
  return func(w http.ResponseWriter, r *http.Request) {
    ctx := NewContext(w, r)
    startTime := time.Now()
    values := []reflect.Value{reflect.ValueOf(ctx)}
    if t.NumIn() == 2 {
      b, err := ioutil.ReadAll(r.Body)
      if err != nil {
        ctx.ErrorInternal(err)
        return
      }
      value := reflect.New(t.In(1))
      if err = json.Unmarshal(b, value.Interface()); err != nil {
        ctx.ErrorBadRequest("Invalid json payload.", nil)
        return
      }
      values = append(values, value.Elem())
    }
    f.Call(values)
    ctx.Log("http_method", ctx.Request.Method)
    ctx.Log("http_path", ctx.Request.RequestURI)
    ctx.Log("http_status", ctx.Response.StatusCode())
    ctx.Log("remote_addr", ctx.Request.RemoteAddr)
    ctx.Log("duration", time.Now().Sub(startTime))
    logrus.WithFields(ctx.LogFields()).Info("canonical-log-line")
  }
}

func (r *Response) WriteHeader(statusCode int) {
  r.statusCode = statusCode
  r.ResponseWriter.WriteHeader(statusCode)
}

func (r *Response) StatusCode() int {
  return r.statusCode
}
