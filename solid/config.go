package solid

var serverConfig = NewServerConfig()
var settings = NewSettingsConfig()

type ServerConfigStruct struct {
	port int
	mainStruct SolidRoute
}

func NewServerConfig() *ServerConfigStruct {
	return &ServerConfigStruct{
		port: 8000,
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
	multipartFormMaxMemory int64
}

func NewSettingsConfig() *SettingsConfigStruct {
	return &SettingsConfigStruct{
		multipartFormMaxMemory: 32 << 20,
	}
}

func (s *SettingsConfigStruct) SetMultipartFormMaxMemory(maxMemory int64) {
	s.multipartFormMaxMemory = maxMemory
}

func (s *SettingsConfigStruct) GetMultipartFormMaxMemory() int64 {
	return s.multipartFormMaxMemory
}

func GetSettingsConfig() *SettingsConfigStruct {
	return settings
}