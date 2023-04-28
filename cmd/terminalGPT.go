package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"terminalGPT/config"
	"terminalGPT/internal/app/flags"
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

	query, newApiKey := flags.Parse()
	if query == "" && newApiKey == "" {
		userInterface := ui.NewUserInterface(gptModel)
		if err := userInterface.Start(); err != nil {
			log.Fatal(err)
		}
	} else {
		if newApiKey != "" {
			if err := configService.SetApiKey(newApiKey); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Api key is successfully modified")
			return
		}
		if query != "" {
			response, err := gptModel.GenerateText(query)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(response)
			return
		}
	}
}
