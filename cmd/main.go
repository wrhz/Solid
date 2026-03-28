package main

import (
	"fmt"
	"net/http"
	"solid/middleware"
	"solid/route"
	"solid/solid"
	"strconv"
)

func main() {
	server, err := solid.GetServer()
	if err != nil {
		fmt.Println("Failed to read server config:", err)
		return
	}

	serve := http.NewServeMux()

	route.Routes()

	middleware.Middleware()

	for path, callFunc := range solid.GetRoutes() {
		handler := solid.Chain(http.HandlerFunc(callFunc));
		serve.Handle("GET " + path, handler)
	}

	fmt.Println("Server starting on port:", server.Port)
    if err := http.ListenAndServe(":" + strconv.Itoa(server.Port), serve); err != nil {
        fmt.Println("Server failed:", err)
    }
}