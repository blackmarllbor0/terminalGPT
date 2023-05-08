package config

import (
	"fmt"
	"log"
	"os"
	yamlI "terminalGPT/pkg/dir/yaml/interfaces"

	configI "terminalGPT/config/interfaces"
)

const (
	apiKeyConfigName = "api-key"
)

type Config struct {
	configReader configI.ConfigReader
	yamlSetter   yamlI.IYAML
}

func NewConfig(configReader configI.ConfigReader, yamlSetter yamlI.IYAML) *Config {
	return &Config{
		configReader: configReader,
		yamlSetter:   yamlSetter,
	}
}

func (c *Config) LoadConfig(pathToConfigFile, configFileType, configFileName string) error {
	c.configReader.SetConfigName(configFileName)
	c.configReader.SetConfigType(configFileType)
	c.configReader.AddConfigPath(pathToConfigFile)
	if err := c.configReader.ReadInConfig(); err != nil {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", pathToConfigFile, configFileName)); err != nil {
			newConfigFile, err := os.Create(fmt.Sprintf("%s/%s.%s", pathToConfigFile, configFileName, configFileType))
			if err != nil {
				return fmt.Errorf("failed to create config file")
			}

			if err := newConfigFile.Close(); err != nil {
				return fmt.Errorf("failed to close config file: %s", err)
			}

			return nil
		}

		return err
	}

	return nil
}

func (c *Config) GetApiKey() (string, error) {
	apiKey := c.configReader.GetString(apiKeyConfigName)
	if apiKey == "" {
		return "", fmt.Errorf("api-key is empty")
	}

	return apiKey, nil
}

func (c *Config) SetApiKey(apiKey string) error {
	if apiKey == "" {
		return fmt.Errorf("api-key cannot be an empty string")
	}

	if err := c.yamlSetter.SetValuesByKeys(map[string]interface{}{apiKeyConfigName: apiKey}); err != nil {
		return err
	}

	log.Println("api-key successfully modified")

	return nil
}
