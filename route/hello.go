package route

import "solid/solid"

type HelloURLArgs struct {
	Name []string `path:"name"`
	Age int `path:"age"`
}

type Hello struct {}

func (h *Hello) RegisterRoute(r *solid.RouteStruct) {
	r.Get("/hello", h.helloFunc)
}

func (h *Hello) RegisterMiddleware(r *solid.RouteStruct) {
	
}

func (h *Hello) helloFunc(c *solid.Context) {
	var args HelloURLArgs

	c.BindQuery(&args)

	solid.JsonResponse(c, args, 200)
}

func NewHello() *Hello {
	return &Hello{}
}