package main

import (
	"net/http"
	"gopkg.in/macaron.v1"

	log "github.com/cihub/seelog"
)

func main() {
	m := macaron.Classic()
	m.Get("/", myHandler)

	log.Info("Server is running...")
	log.Info(http.ListenAndServe("0.0.0.0:4000", m))
}

func myHandler(ctx *macaron.Context) string {
	return "the request path is: " + ctx.Req.RequestURI
}