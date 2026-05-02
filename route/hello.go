package route

import (
	"fmt"
	"solid/solid"
)

type CookieData struct {
	Message int `cookie:"message"`
}

type Hello struct{}

func (h *Hello) RegisterRoute(r *solid.RouteStruct) {
	r.Get("/hello", h.helloFuncGet)
}

func (h *Hello) RegisterMiddleware(r *solid.RouteStruct) {

}

func (h *Hello) helloFuncGet(c *solid.Context) {
	var cookie CookieData
	if err := c.BindCookie(&cookie); err == nil {
		fmt.Println("Cookie message:", cookie.Message)

		cookie.Message++
		if err := c.SaveCookie(&cookie, &solid.CookieOption{}); err != nil {
			fmt.Println("Failed to save cookie:", err)
		}
	}

	solid.ViewResponse(c, "index.html", 200)
}

func NewHello() *Hello {
	return &Hello{}
}