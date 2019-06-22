package main

import (
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

func handler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(200)
	ctx.SetBody([]byte("ok"))
}

type fasthttpLogger struct{}

func (f fasthttpLogger) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func main() {
	server := &fasthttp.Server{
		Logger:  &fasthttpLogger{},
		Handler: handler,
		ErrorHandler: func(ctx *fasthttp.RequestCtx, err error) {
			log.Printf("%v", err)
		},
		Concurrency:        256 * 1024,
		DisableKeepalive:   false,
		ReadTimeout:        10 * time.Second,
		WriteTimeout:       10 * time.Second,
		IdleTimeout:        1 * time.Second,
		ReadBufferSize:     4096,
		WriteBufferSize:    4096,
		MaxConnsPerIP:      256,
		MaxRequestsPerConn: 0,
		MaxRequestBodySize: 4 * 1024 * 1024,
		//
		LogAllErrors:                  false,
		DisableHeaderNamesNormalizing: false,
	}

	log.Println("Start fasthttp server on :8080")
	panic(server.ListenAndServe(":8080"))
}
