package server

import (
  "pb/web"
  "pb/mustache"
)

func routes() {
  web.Post("/", paste) //Defined in paste.go
  web.Get("/", index)
}

func index(ctx *web.Context) string {
  return mustache.RenderFile("templates/index.html")
}
