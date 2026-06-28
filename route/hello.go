package route

import "github.com/wrhz/Solid"

type Hello struct{}

func NewHello() *Hello {
	return &Hello{}
}

func (h *Hello) RegisterRoute(r *solid.RouteStruct) {
	r.Get("/hello", h.helloGet)
}

func (h *Hello) Init(r *solid.RouteStruct) {
	
}	

func (h *Hello) RegisterMiddleware(r *solid.RouteStruct) {
	
}

func (h *Hello) ServerStart() {

}

func (h *Hello) ServerEnd() {

}

func (h *Hello) helloGet(c *solid.Context) error {
	return solid.HtmlViewResponse(c, "index", 200)
}
