package main

import (
	"encoding/json"
	"os"
)

type Server struct {
	Port int `json:"port"`
}

func GetServer() (*Server, error) {
	var data, err = os.ReadFile("config/server.json")
	if err != nil {
		return nil, err
	}
	var server Server
	err = json.Unmarshal(data, &server)
	if err != nil {
		return nil, err
	}
	return &server, nil
}
