package config

import (
	"encoding/json"
	"os"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New)

type Config struct {
	Database database `json:"database"`
}

type database struct {
	Dsn string `json:"dsn"`
}

func New() (*Config, error) {
	fp, err := os.Open("config/app.json")
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var cfg Config
	if err := json.NewDecoder(fp).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
