package main

import (
	"fmt"
	"net/http"
	"solid/config"
	"solid/solid"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	server := config.ServerConfig()

	serve := mux.NewRouter()

	route := solid.NewRoute()

	server.MainStruct.RegisterMiddleware(route)

	server.MainStruct.RegisterRoute(route)

	for path, callFunc := range solid.GetRoutes() {
		serve.HandleFunc(path, callFunc).Methods("GET")
	}

	fmt.Println("Server starting on port:", server.Port)
    if err := http.ListenAndServe(":" + strconv.Itoa(server.Port), serve); err != nil {
        fmt.Println("Server failed:", err)
    }
}