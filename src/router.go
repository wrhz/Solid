package src

import "net/http"

type Router struct {}

var getRoutes = map[string]func(w http.ResponseWriter, r *http.Request) {}

func NewRouter() *Router {
	return &Router{}
}

func GetRoutes() map[string]func(w http.ResponseWriter, r *http.Request) {
	return getRoutes
}

func (r* Router) Get(path string, callFunc func(w http.ResponseWriter, r *http.Request)) {
	getRoutes[path] = callFunc
}
