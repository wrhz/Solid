package main

import (
	"fmt"
	"net/http"
	"solid/config"
	"solid/solid"
	"strconv"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
)

func main() {
	config.ServerConfig()

	config.SettingsConfig()

	serverConfig := solid.GetServerConfig()

	serve := mux.NewRouter()

	route := solid.NewRoute()

	serverConfig.GetMainStruct().Init(route)

	serverConfig.GetMainStruct().RegisterMiddleware(route)

	serverConfig.GetMainStruct().RegisterRoute(route)

	for path, callFunc := range solid.GetRoutes() {
		serve.HandleFunc(path, callFunc).Methods("GET")
	}

	for path, callFunc := range solid.PostRoutes() {
		serve.HandleFunc(path, callFunc).Methods("POST")
	}

	serve.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	)

	wrappedHandler := gziphandler.GzipHandler(serve)

	fmt.Println("Server starting on port:", serverConfig.GetPort())
	if err := http.ListenAndServe(":"+strconv.Itoa(serverConfig.GetPort()), wrappedHandler); err != nil {
		fmt.Println("Server failed:", err)
	}
}
