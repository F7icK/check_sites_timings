package main

import (
	"flag"
	"os"

	"github.com/F7icK/check_sites_timings/internal/application"
	"github.com/F7icK/check_sites_timings/internal/check_sites_timings/types/config"
	"github.com/F7icK/check_sites_timings/pkg/customlog"
	"gopkg.in/yaml.v2"
)

func main() {
	configPath := new(string)

	flag.StringVar(configPath, "config-path", "./config/config.yaml", "specify path to yaml")
	flag.Parse()

	customlog.Init()

	configFile, err := os.Open(*configPath)
	if err != nil {
		customlog.Error.Println(err)
		return
	}

	cfg := config.Config{}
	if err = yaml.NewDecoder(configFile).Decode(&cfg); err != nil {
		customlog.Error.Println(err)
		return
	}

	application.NewApplication(&cfg)
}
