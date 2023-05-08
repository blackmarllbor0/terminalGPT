package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"terminalGPT/config"
	"terminalGPT/internal/app/flags"
	"terminalGPT/internal/app/repository/mongo"
	"terminalGPT/internal/app/repository/mongo/models"
	"terminalGPT/internal/pkg/api/gpt3dot5t0urbo"
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

	ctx := context.Background()
	DBService := mongo.NewClient(configReaderService, ctx)
	if err := DBService.Connection(); err != nil {
		log.Fatal(err)
	}
	defer DBService.Disconnect()

	chatsService := models.NewChats(DBService.GetCollectionByName(mongo.CHATS), ctx)
	gptModel := gpt3dot5t0urbo.NewGPT3(configReaderService, chatsService)

	query, newAPIKey := flags.Parse()
	if query == "" && newAPIKey == "" {
		userInterface := ui.NewUserInterface(gptModel)
		if err := userInterface.Start(); err != nil {
			log.Println(err)
			return
		}
	} else {
		if newAPIKey != "" {
			if err := configService.SetApiKey(newAPIKey); err != nil {
				log.Println(err)
				return
			}

			fmt.Println("Api key is successfully modified")
			return
		}
		if query != "" {
			response, err := gptModel.GenerateText(query, primitive.ObjectID{})
			if err != nil {
				log.Println(err)
				return
			}

			fmt.Println(response)
			return
		}
	}
}
