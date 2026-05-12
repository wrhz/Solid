package route

import (
	"fmt"
	"solid/solid"
)

type SessionData struct {
	Message string `session:"message"`
}

type Hello struct{}

func (h *Hello) RegisterRoute(r *solid.RouteStruct) {
	r.Get("/setSession", h.setSession)
	r.Get("/getSession", h.getSession)
}

func (h *Hello) RegisterMiddleware(r *solid.RouteStruct) {

}

func (h *Hello) setSession(c *solid.Context) {
	var data = &SessionData{
		Message: "Hello Solid",
	}

	c.SaveSession(data, "hello-session", &solid.SessionOptions{})

	solid.StringResponse(c, "Session set successfully", 200)
}

func (h *Hello) getSession(c *solid.Context) {
	var data SessionData

	if err := c.BindSession(&data, "hello-session"); err != nil {
		c.Error(500, err)
		return
	}

	fmt.Println("Session message:", data.Message)

	solid.StringResponse(c, data.Message, 200)
}

func NewHello() *Hello {
	return &Hello{}
}