package httpx

import "net/http"

type Response struct {
  http.ResponseWriter
  statusCode int
}

func (r *Response) WriteHeader(statusCode int) {
  r.statusCode = statusCode
  r.ResponseWriter.WriteHeader(statusCode)
}

func (r *Response) StatusCode() int {
  return r.statusCode
}
