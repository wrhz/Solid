package config

import "solid/solid"

func SettingsConfig() {
	settings := solid.GetSettingsConfig()
	
	settings.SetMultipartFormMaxMemory(64 << 20)
}