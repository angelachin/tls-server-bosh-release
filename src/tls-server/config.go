package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ServerCACertFile          string `json:"ca_cert_file" validate:"nonzero"`
	ServerCertFile            string `json:"server_cert_file" validate:"nonzero"`
	ServerKeyFile             string `json:"server_key_file" validate:"nonzero"`
	ListenPort                  int    `json:"listen_port" validate:"min=1"`
}

func LoadConfig(filePath string) (Config, error) {
	var cfg Config
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return cfg, fmt.Errorf("reading file %s: %s", filePath, err)
	}

	err = json.Unmarshal(contents, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("unmarshaling contents: %s", err)
	}
	return cfg, nil
}
