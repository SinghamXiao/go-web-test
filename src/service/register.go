package service

import (
	"gopkg.in/macaron.v1"
	"io/ioutil"
	"fmt"
	. "domain"
)

func RegisterRouter(m *macaron.Macaron) {
	m.Post("/auth/register", registerHandler)
}

func registerHandler(ctx *macaron.Context) []byte {
	body, err := ioutil.ReadAll(ctx.Req.Request.Body)
	if err != nil {
		fmt.Println(string(body))
	}

	registerInfo := DecodeRegisterInfo(body)
	if registerInfo != nil {
		fmt.Println("Username: ", registerInfo.Username)
	}

	return body
}
