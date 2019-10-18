package middleware

import (
  "github.com/remicaumette/zaap.sh/pkg/util/httpx"
  "github.com/sirupsen/logrus"
  "time"
)

func CanonicalLog(ctx *httpx.Context, next httpx.NextFunc) {
  startTime := time.Now()
  next()
  logrus.
    WithField("duration", time.Now().Sub(startTime)).
    WithField("http_method", ctx.Request.Method).
    WithField("http_path", ctx.Request.RequestURI).
    WithField("http_status", ctx.Response.StatusCode()).
    Info("canonical-log-line")
}

//http_method=POST http_path=/v1/charges http_status=200 key_id=mk_123 permissions_used=account_write
//rate_allowed=true rate_quota=100 rate_remaining=99 request_id=req_123 team=acquiring user_id=usr_123
