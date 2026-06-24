package config

import (
	"solid/model"
	"solid/solid"

	_ "github.com/mattn/go-sqlite3"
)

func DatabaseConfig() {
	database := solid.GetDatabaseConfig()

	database.SetXormDriverName("sqlite3")
	database.SetXormDataSourceName("./example.db")

	database.RegisterXormModels(&model.User{})
}