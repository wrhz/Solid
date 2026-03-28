package solid

import "net/http"

type Middleware struct{}

var middlewares []func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func Chain(handler http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		next := handler
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			middlewares[i](w, r, next.ServeHTTP)
		})
	}

	return handler
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (c *Middleware) Use(middleware func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {
	middlewares = append(middlewares, middleware)
}