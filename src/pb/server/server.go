package server

import (
	"fmt"
	"pb/web"
)

func Start() {
	Log("Initializing Routes...")
	routes() //defined in routes.go
	Log("Starting server...")
	web.Run("0.0.0.0:8080")
}

func Log(m interface{}) {
	DEBUG := true
	if DEBUG {
		fmt.Println(m)
	}
}
