package main

import (
	"fmt"
	"net/http"
	"solid/config"
	"solid/solid"
	"strconv"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleRoutes(serve *mux.Router) {
	for path, callFunc := range solid.GetRoutes() {
		serve.Handle(path, gziphandler.GzipHandler(http.HandlerFunc(callFunc))).Methods("GET")
	}

	for path, callFunc := range solid.PostRoutes() {
		serve.Handle(path, gziphandler.GzipHandler(http.HandlerFunc(callFunc))).Methods("POST")
	}

	for path, callFunc := range solid.PatchRoutes() {
		serve.Handle(path, gziphandler.GzipHandler(http.HandlerFunc(callFunc))).Methods("PATCH")
	}

	for path, callFunc := range solid.DeleteRoutes() {
		serve.Handle(path, gziphandler.GzipHandler(http.HandlerFunc(callFunc))).Methods("DELETE")
	}

	for path, callFunc := range solid.PutRoutes() {
		serve.Handle(path, gziphandler.GzipHandler(http.HandlerFunc(callFunc))).Methods("PUT")
	}

	for path, callFunc := range solid.OptionsRoutes() {
		serve.Handle(path, gziphandler.GzipHandler(http.HandlerFunc(callFunc))).Methods("OPTIONS")
	}

	for path, callFunc := range solid.HeadRoutes() {
		serve.Handle(path, gziphandler.GzipHandler(http.HandlerFunc(callFunc))).Methods("HEAD")
	}

	for path, callFunc := range solid.WebsocketRoutes() {
		serve.Handle(path, http.HandlerFunc(callFunc))
	}
}

func main() {
	config.ServerConfig()

	config.SettingsConfig()

	serverConfig := solid.GetServerConfig()

	serve := mux.NewRouter()

	route := solid.NewRoute()

	serverConfig.GetMainStruct().Init(route)

	serverConfig.GetMainStruct().RegisterMiddleware(route)

	serverConfig.GetMainStruct().RegisterRoute(route)

	handleRoutes(serve)

	serve.PathPrefix("/static/").Handler(
		gziphandler.GzipHandler(
			http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
		),
	)

	fmt.Println("Server starting on port:", serverConfig.GetPort())

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(serverConfig.GetPort()),
		Handler: serve,
	}

	if tlsConfig := solid.GetServerConfig().GetTLSConfig(); tlsConfig != nil {
		server.TLSConfig = tlsConfig
	}

	if certFile := solid.GetServerConfig().GetTLSCertFile(); certFile != "" {
		keyFile := solid.GetServerConfig().GetTLSKeyFile()

		if err := server.ListenAndServeTLS("./certs/" + certFile, "./certs/" + keyFile); err != nil {
			fmt.Println("Server failed:", err)
		}
		return
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server failed:", err)
	}
}
