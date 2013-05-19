package server

import (
  "pb/web"
  "pb/mustache"
)

func routes() {
  web.Post("/", paste)
  web.Get("/", index)
}

func paste(ctx *web.Context) string {
  return "Done."
}

func index(ctx *web.Context) string {
  return mustache.RenderFile("templates/index.html")
}
