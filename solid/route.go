package solid

import (
	// "fmt"
	"fmt"
	"net/http"
	"strings"
	"time"
	// "strings"
	// "time"
)

type RouteStruct struct {
	perfix      string
	middlewares []func(http.Handler) http.Handler
}

type SolidRoute interface {
	RegisterRoute(*RouteStruct)
	RegisterMiddleware(*RouteStruct)
}

var getRoutes = map[string]func(w http.ResponseWriter, r *http.Request){}
var postRoutes = map[string]func(w http.ResponseWriter, r *http.Request){}

func GetRoutes() map[string]func(w http.ResponseWriter, r *http.Request) {
	return getRoutes
}

func PostRoutes() map[string]func(w http.ResponseWriter, r *http.Request) {
	return postRoutes
}

func NewRoute() *RouteStruct {
	return &RouteStruct{perfix: ""}
}

func (r *RouteStruct) Get(path string, callFunc func(c *Context)) {
	getRoutes[r.perfix+path] = r.chain(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		callFunc(&Context{Writer: w, Request: req})
	})).ServeHTTP
}

func (r *RouteStruct) Post(path string, callFunc func(c *Context)) {
	postRoutes[r.perfix+path] = r.chain(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		callFunc(&Context{Writer: w, Request: req})
	})).ServeHTTP
}

func (r *RouteStruct) Group(prefix string, callStruct SolidRoute) {
	route := &RouteStruct{perfix: r.perfix + prefix, middlewares: r.middlewares}
	callStruct.RegisterMiddleware(route)
	callStruct.RegisterRoute(route)
}

func (r *RouteStruct) Use(middleware func(c *Context, next http.HandlerFunc)) {
	r.middlewares = append(r.middlewares, func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			middleware(&Context{Writer: w, Request: r}, next.ServeHTTP)
		})
	})
}

func (r *RouteStruct) chain(handler http.Handler) http.Handler {
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}

	originalHandler := handler

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func ()  {
			if err := recover(); err != nil {
				fmt.Printf("Panic recovered: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		settings := GetSettingsConfig()

		if strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
			err := r.ParseMultipartForm(settings.GetMultipartFormMaxMemory())

			if err != nil {
				fmt.Println("Error parsing multipart form:", err)
				http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
				return
			}
		}

		timeStart := time.Now()

		originalHandler.ServeHTTP(w, r)

		fmt.Printf("[%s] %s %s ... %v\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, time.Since(timeStart))
	})
}
