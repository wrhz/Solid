package config

import "solid/solid"

func WebSocketConfig() *solid.WebSocketConfigStruct {
	websocket := solid.NewWebSocketConfig()

	return websocket
}