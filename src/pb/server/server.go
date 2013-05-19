package server

import (
  "pb/web"
)

func Start() {
  routes() //defined in routes.go
  web.Run("0.0.0.0:8080")
}
