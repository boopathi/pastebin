package server

import (
  "pb/web"
  "pb/mustache"
  "pb/redis"
)

func paste(ctx *web.Context) string {
  var store redis.Client
  err := store.Set("hello", []byte("world"))
  if err != nil {
    return mustache.Render("Database Error.")
  }
  ctx.SetHeader("X-Powered-By","web.go",true)
  ctx.SetHeader("Connection", "close",true)
  return mustache.Render("Done")
}
