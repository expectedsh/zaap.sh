package server

import (
  "github.com/remicaumette/zaap.sh/pkg/util/httpx"
)

func corsMiddleware() httpx.MiddlewareFunc {
  return func(ctx *httpx.Context, next httpx.NextFunc) {
    ctx.
      Header("Access-Control-Allow-Origin", "*").
      Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE").
      Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
    next()
  }
}
