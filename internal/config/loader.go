package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Query []Query `yaml:"query"`
}

// Query represents the query config map.
type Query struct {
	Cmd       string `yaml:"cmd"`
	URL       string `yaml:"url"`
	SearchStr string `yaml:"search_str"`
}

// QueryConfig represents the config used in query package.
var QueryConfig = make(map[string]Query)

func loadQueryConfig(cf *os.File) {
	decoder := yaml.NewDecoder(cf)
	var configTemp config
	err := decoder.Decode(&configTemp)
	if err != nil {
		log.Fatal("Failed to decode config file.")
	}

	// Transform the query config into a map using command string as key.
	for _, qc := range configTemp.Query {
		QueryConfig[qc.Cmd] = qc
	}
}

// Load loads all the configs used to run the program.
func Load() {
	wd, _ := os.Getwd()
	cf, err := os.Open(wd + "../config/config.yml")
	if err != nil {
		log.Fatal("Cannot find config file.")
	}
	defer cf.Close()

	loadQueryConfig(cf)
}
