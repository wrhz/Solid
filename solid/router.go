package solid

import "net/http"

type Router struct {}

var getRoutes = map[string]func(w http.ResponseWriter, r *http.Request) {}

func GetRoutes() map[string]func(w http.ResponseWriter, r *http.Request) {
	return getRoutes
}

func NewRouter() *Router {
	return &Router{}
}

func (r* Router) Get(path string, callFunc func(c *Context)) {
	getRoutes[path] = func(w http.ResponseWriter, r *http.Request) {
		callFunc(&Context{w: w, r: r})
	}
}
