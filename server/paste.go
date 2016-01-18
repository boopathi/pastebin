package server

import (
	"crypto/rand"
	"github.com/garyburd/redigo/redis"
	"github.com/hoisie/mustache"
	"github.com/hoisie/web"
	"io"
	"regexp"
)

func rands(l int) string {
	chars := make([]byte, 62)
	chars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]byte, l)
	r := make([]byte, l+(l/4))
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, r); err != nil {
			panic("error reading from random source: " + err.Error())
		}
		for _, c := range r {
			if c >= maxrb {
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == l {
				return string(b)
			}
		}
	}
	panic("unreachable")
}

func paste(ctx *web.Context) string {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		Log(err)
		return mustache.Render("Cannot connect to DB \r\n")
	}
	defer c.Close()
	rstr := rands(5)
	if val, ok := ctx.Params["name"]; ok {
		if matched, _ := regexp.MatchString("[a-zA-Z0-9]+", val); matched {
			rstr = val
		}
	}
	y := "paste_" + rstr
	c.Do("SET", y, ctx.Params["paste"])
	sname := "/" //"http://" + ctx.Server.Config.Addr + ":" + string(ctx.Server.Config.Port) + "/"
	return mustache.Render(sname + rstr + "\n")
}

func getPaste(ctx *web.Context, uri string) string {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		Log(err)
		return mustache.Render("Cannot connect to DB \r\n")
	}
	defer c.Close()
	y := "PASTE_DOES_NOT_EXIST"
	if matched, _ := regexp.MatchString("[a-zA-Z0-9]+", uri); matched {
		y = "paste_" + uri
	}
	x, err := redis.String(c.Do("GET", y))
	ctx.SetHeader("Content-Type", "text/plain", true)
	ctx.SetHeader("X-Powered-By", "web.go", true)
	ctx.SetHeader("Connection", "close", true)
	return mustache.Render(x)
}
