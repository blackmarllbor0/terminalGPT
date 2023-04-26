package yaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type YAML struct {
	path string
}

func NewYAML(path string) *YAML {
	return &YAML{
		path: path,
	}
}

func (y YAML) SetValuesByKeys(keyAndValues map[string]interface{}) error {
	data, err := os.ReadFile(y.path)
	if err != nil {
		return fmt.Errorf("the file could not be opened: %v", err)
	}

	structWithKeys := make(map[string]interface{})
	if err := yaml.Unmarshal(data, &structWithKeys); err != nil {
		return fmt.Errorf("the file could not be parsed: %v", err)
	}

	for key, value := range keyAndValues {
		structWithKeys[key] = value
	}

	data, err = yaml.Marshal(&structWithKeys)
	if err != nil {
		return fmt.Errorf("the file could not be parsed: %v", err)
	}

	if err := os.WriteFile(y.path, data, 0644); err != nil {
		return fmt.Errorf("the file could not be written: %v", err)
	}

	return nil
}
