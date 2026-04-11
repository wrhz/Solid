package solid

type ServerConfigStruct struct {
	Port int
	MainStruct SolidRoute
}

func NewServerConfig() *ServerConfigStruct {
	return &ServerConfigStruct{
		Port: 8000,
		MainStruct: nil,
	}
}

func (s *ServerConfigStruct) SetPort(port int) {
	s.Port = port
}

func (s *ServerConfigStruct) SetMainStruct(mainStruct SolidRoute) {
	s.MainStruct = mainStruct
}