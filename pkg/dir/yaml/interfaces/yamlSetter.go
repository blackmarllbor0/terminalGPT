package interfaces

type IYAML interface {
	SetValuesByKeys(keyAndValues map[string]interface{}) error
}
