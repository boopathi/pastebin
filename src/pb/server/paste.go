package server

import (
  "pb/web"
  "pb/mustache"
  "pb/redis"
)

func paste(ctx *web.Context) string {
  var store redis.Client
  ctx.SetHeader("X-Powered-By","web.go",true)
  ctx.SetHeader("Connection", "close",true)
  return "URI: 82j8928"
}
