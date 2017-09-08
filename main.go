package main

import (
	"net/http"

	. "service"

	"gopkg.in/macaron.v1"
	log "github.com/cihub/seelog"
)

var m *macaron.Macaron

func init() {
	m = macaron.Classic()
}

func main() {
	log.Warn("Golang Web Test Server")

	RouterBinding(m) // 路由绑定函数

	err := http.ListenAndServe(":9090", m) //设置监听的端口
	if err != nil {
		log.Error("ListenAndServe: ", err)
	}
}
