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

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(serverConfig.GetPort()),
		Handler: wrappedHandler,
	}

	if tlsConfig := solid.GetServerConfig().GetTLSConfig(); tlsConfig != nil {
		server.TLSConfig = tlsConfig
	}

	if certFile := solid.GetServerConfig().GetTLSCertFile(); certFile != "" {
		keyFile := solid.GetServerConfig().GetTLSKeyFile()

		if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
			fmt.Println("Server failed:", err)
		}
		return
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server failed:", err)
	}
}
