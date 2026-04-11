package main

import (
	"fmt"
	"net/http"
	"solid/config"
	"solid/solid"
	"strconv"
)

func main() {
	server := config.ServerConfig()

	serve := http.NewServeMux()

	route := solid.NewRoute()

	server.MainStruct.RegisterMiddleware(route)

	server.MainStruct.RegisterRoute(route)

	for path, callFunc := range solid.GetRoutes() {
		serve.Handle("GET " + path, http.HandlerFunc(callFunc))
	}

	fmt.Println("Server starting on port:", server.Port)
    if err := http.ListenAndServe(":" + strconv.Itoa(server.Port), serve); err != nil {
        fmt.Println("Server failed:", err)
    }
}