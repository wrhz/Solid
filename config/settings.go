package config

import (
	"solid/solid"

	"github.com/gorilla/sessions"
)

func SettingsConfig() {
	settings := solid.GetSettingsConfig()

	settings.SetMaxBytesMemory(1 << 20)

	secret := []byte("your-sessions-key-must-16|24|32-bytes-long!!!")

	settings.SetSessionStore(sessions.NewCookieStore(secret))
}