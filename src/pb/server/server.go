package server

import (
  "pb/web"
)

func Start() {
  routes()
  web.Run("0.0.0.0:8080")
}
