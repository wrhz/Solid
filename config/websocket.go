package config

import "solid/solid"

func WebSocketConfig() *solid.WebSocketConfigStruct {
	websocket := solid.GetWebSocketConfig()

	return websocket
}