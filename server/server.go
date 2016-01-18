package server

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/hoisie/web"
	"strconv"
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
	var port int
	flag.IntVar(&port, "p", 8080, "Port number")
	flag.Parse()
	web.Run("0.0.0.0:" + strconv.Itoa(port))
}

func Log(m interface{}) {
	DEBUG := true
	if DEBUG {
		fmt.Println(m)
	}
}
