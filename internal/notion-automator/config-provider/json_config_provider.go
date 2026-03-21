package configprovider

import (
	"encoding/json"
	"os"
)

type JsonConfigProvider struct{}

func (jsonConfigProvider JsonConfigProvider) GetConfig(filepath string) ConfigMain {
	configRaw, err := os.ReadFile(filepath)
	if err != nil {
		panic("Cannot read the config file: " + err.Error())
	}

	var config ConfigMain
	err = json.Unmarshal(configRaw, &config)
	if err != nil {
		panic("Cannot unmarshal the config file: " + err.Error())
	}

	return config
}
