package server

import (
	"fmt"
	"pb/redigo/redis"
	"pb/web"
)

func Start() {
	Log("Initializing Routes...")
	routes() //defined in routes.go
	Log("Connecting to Redis Store")
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		Log("Couldn't connect to Redis Store...")
		return
	}
	defer c.Close()
	Log("Starting server...")
	web.Run("0.0.0.0:8080")
}

func Log(m interface{}) {
	DEBUG := true
	if DEBUG {
		fmt.Println(m)
	}
}
