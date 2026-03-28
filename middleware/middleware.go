package middleware

import (
	"fmt"
	"net/http"
	"solid/solid"
	"time"
)

func Middleware() {
	var middleware = solid.NewMiddleware()

	middleware.Use(loggerMiddleware)
}

func loggerMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s %s\n", t, r.Method, r.URL.Path)
}