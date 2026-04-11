package solid

import "net/http"

type RouteStruct struct {
	perfix string
	middlewares []func (http.Handler) http.Handler
}

type SolidRoute interface {
	RegisterRoute(*RouteStruct)
	RegisterMiddleware(*RouteStruct)
}

var getRoutes = map[string]func(w http.ResponseWriter, r *http.Request) {}

func GetRoutes() map[string]func(w http.ResponseWriter, r *http.Request) {
	return getRoutes
}

func NewRoute() *RouteStruct {
	return &RouteStruct{ perfix: "" }
}

func (r *RouteStruct) Get(path string, callFunc func(c *Context)) {
	getRoutes[r.perfix + path] = r.chain(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		callFunc(&Context{Writer: w, Request: req})
	})).ServeHTTP
}

func (r* RouteStruct) Group(prefix string, callStruct SolidRoute) {
	route := &RouteStruct{ perfix: r.perfix + prefix, middlewares: r.middlewares }
	callStruct.RegisterMiddleware(route)
	callStruct.RegisterRoute(route)
}

func (r *RouteStruct) Use(middleware func(c *Context, next http.HandlerFunc)) {
	r.middlewares = append(r.middlewares, func(next http.Handler) http.Handler {
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			middleware(&Context{Writer: w, Request: r}, next.ServeHTTP)
		})
	})
}

func (r *RouteStruct) chain(handler http.Handler) http.Handler {
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}

	return handler
}
