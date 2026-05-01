package route

import (
	"fmt"
	"solid/solid"
)

type Hello struct{}

func (h *Hello) RegisterRoute(r *solid.RouteStruct) {
	r.Get("/hello", h.helloFuncGet)
}

func (h *Hello) RegisterMiddleware(r *solid.RouteStruct) {

}

func (h *Hello) helloFuncGet(c *solid.Context) {
	session, err := c.Session("test-session")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	message := session.GetSession("message")

	if message == nil {
		err := session.SetSession("message", "Hello, World!")
		if err != nil {
			fmt.Println("Error setting session:", err)
			return
		}
		fmt.Println("Session message: nil")
	} else {
		fmt.Println("Session message:", message)
	}
	
	solid.ViewResponse(c, "index.html", 200)
}

func NewHello() *Hello {
	return &Hello{}
}