package service

import "gopkg.in/macaron.v1"

func RouterBinding(m *macaron.Macaron)  {
	rootRouter(m)
	RegisterRouter(m)
	//LoginRouter(m)
}

func rootRouter(m *macaron.Macaron) {
	m.Get("/", rootHandler)
}

func rootHandler(ctx *macaron.Context) string {
	return "Welcome! Just For Testing!"
}