package main

import (
	"github.com/spf13/viper"
	"log"
	"terminalGPT/config"
	"terminalGPT/internal/pkg/api/GPT3"
	"terminalGPT/pkg/dir/yaml"
	"terminalGPT/pkg/ui"
)

func main() {
	configReaderService := viper.New()
	configSetterService := yaml.NewYAML("config/config.yml")
	configService := config.NewConfig(configReaderService, configSetterService)
	if err := configService.LoadConfig("config", "yml", "config"); err != nil {
		log.Fatal(err)
	}

	gptModel := GPT3.NewGPT3(configReaderService)

	userInterface := ui.NewUserInterface(gptModel)
	if err := userInterface.Start(); err != nil {
		log.Fatal(err)
	}
}
