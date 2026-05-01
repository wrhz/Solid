package solid

import "github.com/gorilla/sessions"

var serverConfig = NewServerConfig()
var settings = NewSettingsConfig()

type ServerConfigStruct struct {
	port       int
	mainStruct SolidRoute
}

func NewServerConfig() *ServerConfigStruct {
	return &ServerConfigStruct{
		port:       8000,
		mainStruct: nil,
	}
}

func (s *ServerConfigStruct) SetPort(port int) {
	s.port = port
}

func (s *ServerConfigStruct) GetPort() int {
	return s.port
}

func (s *ServerConfigStruct) SetMainStruct(mainStruct SolidRoute) {
	s.mainStruct = mainStruct
}

func (s *ServerConfigStruct) GetMainStruct() SolidRoute {
	return s.mainStruct
}

func GetServerConfig() *ServerConfigStruct {
	return serverConfig
}

type SettingsConfigStruct struct {
	maxBytesMemory         int64
	multipartFormMaxMemory int64

	sessionsPairs []byte
	sessionStore   sessions.Store
}

func NewSettingsConfig() *SettingsConfigStruct {
	return &SettingsConfigStruct{
		maxBytesMemory:         64 << 20,
		multipartFormMaxMemory: 32 << 20,
	}
}

func (s *SettingsConfigStruct) SetMaxBytesMemory(maxBytesMemory int64) {
	s.maxBytesMemory = maxBytesMemory
}

func (s *SettingsConfigStruct) SetMultipartFormMaxMemory(maxMemory int64) {
	s.multipartFormMaxMemory = maxMemory
}

func (s *SettingsConfigStruct) SetSessionsSecret(sessionsPairs ...string) {
	for _, pair := range sessionsPairs {
		s.sessionsPairs = append(s.sessionsPairs, []byte(pair)...)
	}
}

func (s *SettingsConfigStruct) SetSessionStore(sessionStore sessions.Store) {
	s.sessionStore = sessionStore
}

func (s *SettingsConfigStruct) GetMaxBytesMemory() int64 {
	return s.maxBytesMemory
}

func (s *SettingsConfigStruct) GetMultipartFormMaxMemory() int64 {
	return s.multipartFormMaxMemory
}

func (s *SettingsConfigStruct) GetSessionStore() sessions.Store {
	return s.sessionStore
}

func GetSettingsConfig() *SettingsConfigStruct {
	return settings
}