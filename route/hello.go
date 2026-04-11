package route

import "solid/solid"

type Hello struct {}

func (h *Hello) RegisterRoute(r *solid.RouteStruct) {
	r.Get("/hello", h.helloFunc)
}

func (h *Hello) RegisterMiddleware(r *solid.RouteStruct) {
	
}

func (h *Hello) helloFunc(c *solid.Context) {
	solid.StringResponse(c, "Hello World", 200)
}

func NewHello() *Hello {
	return &Hello{}
}