package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

func main() {
	var configPath string
	pflag.StringVarP(&configPath, "config", "c", "config.yaml", "path to config file")
	pflag.Parse()

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Cannot read config: %v", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		log.Fatalf("Cannot parse config: %v", err)
	}

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	fmt.Printf("Server starting on Port: %d: ", cfg.Server.Port)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}

}
