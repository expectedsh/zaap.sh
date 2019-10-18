package server

import (
  "github.com/sirupsen/logrus"
  "net/http"
  "time"
)

func (s *Server) logMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    startTime := time.Now()
    next.ServeHTTP(w, r)
    // alloc_count
    // auth_type
    // database_queries
    // duration
    logrus.
      WithField("duration", time.Now().Sub(startTime)).
      WithField("http_method", r.Method).
      WithField("http_path", r.RequestURI).
      WithField("http_status", w.Header().Get("Status-Code")).
      Info("canonical-log-line")
  })
}



 //http_method=POST http_path=/v1/charges http_status=200 key_id=mk_123 permissions_used=account_write
 //rate_allowed=true rate_quota=100 rate_remaining=99 request_id=req_123 team=acquiring user_id=usr_123
