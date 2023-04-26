package main

import (
	"context"
	"fmt"
	"log"
	"terminalGPT/config"
	"terminalGPT/internal/app/flags"
	"terminalGPT/pkg/dir/yaml"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func main() {
	configReaderService := viper.New()
	yamlSetterService := yaml.NewYAML("config/config.yml")
	configService := config.NewConfig(configReaderService, yamlSetterService)
	if err := configService.LoadConfig("config", "yml", "config"); err != nil {
		log.Fatal(err)
	}

	flagService := flags.NewFlags()
	query, newApiKey := flagService.Parse()
	if newApiKey != "" {
		if err := configService.SetApiKey(newApiKey); err != nil {
			log.Fatal(err)
		}
	}

	apiKey, err := configService.GetApiKey()
	if err != nil {
		log.Fatal(err)
	}

	client := openai.NewClient(apiKey)
	response, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{{
			Role:    openai.ChatMessageRoleUser,
			Content: query,
		}},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Choices[0].Message.Content)
}
