package route

import (
	"solid/solid"
)

type Hello struct{}

func (h *Hello) RegisterRoute(r *solid.RouteStruct) {
	r.Get("/hello", h.helloFuncGet)
}

func (h *Hello) RegisterMiddleware(r *solid.RouteStruct) {

}

func (h *Hello) helloFuncGet(c *solid.Context) {
	solid.ViewResponse(c, "index.html", 200)
}

func NewHello() *Hello {
	return &Hello{}
}