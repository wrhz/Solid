package route

import "solid/solid"

type HelloURLArgs struct {
	Id int `param:"id"`
}

type Hello struct {}

func (h *Hello) RegisterRoute(r *solid.RouteStruct) {
	r.Get("/hello/{id:[0-9]+}", h.helloFunc)
}

func (h *Hello) RegisterMiddleware(r *solid.RouteStruct) {
	
}

func (h *Hello) helloFunc(c *solid.Context) {
	var args HelloURLArgs

	if c.BindParams(&args) != nil {
		solid.JsonResponse(c, map[string]string{"error": "invalid parameters"}, 500)
		return
	}

	solid.JsonResponse(c, args, 200)
}

func NewHello() *Hello {
	return &Hello{}
}