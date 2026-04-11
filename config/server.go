package config

import (
	"solid/route"
	"solid/solid"
)

func ServerConfig() *solid.ServerConfigStruct {
    server := solid.NewServerConfig()

    server.SetPort(8000)

    server.SetMainStruct(route.NewHello())

    return server
}