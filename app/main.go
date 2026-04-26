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
	config.ServerConfig()

	config.SettingsConfig()

	serverConfig := solid.GetServerConfig()

	serve := mux.NewRouter()

	route := solid.NewRoute()

	serverConfig.GetMainStruct().RegisterMiddleware(route)

	serverConfig.GetMainStruct().RegisterRoute(route)

	for path, callFunc := range solid.GetRoutes() {
		serve.HandleFunc(path, callFunc).Methods("GET")
	}

	for path, callFunc := range solid.PostRoutes() {
		serve.HandleFunc(path, callFunc).Methods("POST")
	}

	fmt.Println("Server starting on port:", serverConfig.GetPort())
	if err := http.ListenAndServe(":"+strconv.Itoa(serverConfig.GetPort()), serve); err != nil {
		fmt.Println("Server failed:", err)
	}
}
