package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"terminalGPT/config"
	"terminalGPT/internal/app/flags"
	"terminalGPT/internal/app/repository/mongo"
	"terminalGPT/internal/pkg/api/GPT3dot5Turbo"
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

	DBService := mongo.NewClient(configReaderService)
	if err := DBService.Connection(); err != nil {
		log.Fatal(err)
	}

	defer DBService.Disconnect()

	gptModel := GPT3dot5Turbo.NewGPT3(configReaderService)

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
