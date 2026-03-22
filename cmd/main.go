package main

import (
	"fmt"
	"net/http"
	"solid/routes"
	"solid/src"
	"strconv"
)

func main() {
	server, err := GetServer()
	if err != nil {
		fmt.Println("Failed to read server config:", err)
		return
	}

	serveMux := http.NewServeMux()

	routes.Routes()

	for path, callFunc := range src.GetRoutes() {
		serveMux.HandleFunc("GET " + path, callFunc)
	}

	fmt.Println("Server starting on port:", server.Port)
    if err := http.ListenAndServe(":" + strconv.Itoa(server.Port), serveMux); err != nil {
        fmt.Println("Server failed:", err)
    }
}