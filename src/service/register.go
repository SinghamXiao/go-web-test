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
	fmt.Println(string(body))
	userInfo := Decode(body)
	fmt.Println(userInfo.Username)

	return body
}