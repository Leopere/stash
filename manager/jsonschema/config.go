package jsonschema

import (
	"encoding/json"
	"github.com/stashapp/stash/logger"
	"os"
)

type Config struct {
	Stash string `json:"stash"`
	Metadata string `json:"metadata"`
	// Generated string `json:"generated"` // TODO: Generated directory instead of metadata
	Cache string `json:"cache"`
	Downloads string `json:"downloads"`
}

func LoadConfigFile(file string) *Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		logger.Error(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	parseError := jsonParser.Decode(&config)
	if parseError != nil { panic(parseError) }
	return &config
}