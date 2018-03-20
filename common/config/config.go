package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

var defaultConf = []byte(`
core:
  mongo_uri: localhost:27017 # mongo storage uri
  port: 6767 # app's port
  zipkin_uri: http://192.168.99.100:9411 # zipkin uri
  storage: mongo # use storage type.
`)

// ConfYaml is config structure.
type ConfYaml struct {
	Core SectionCore `yaml:"core"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	MongoURI  string `yaml:"mongo_uri"`
	Port      string `yaml:"port"`
	ZipkinURI string `yaml:"zipkin_uri"`
	Storage   string `yaml:"storage"`
}

// LoadConf load config from file and read in environment variables that match
func LoadConf(confPath string) (ConfYaml, error) {
	var conf ConfYaml
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("gojobs")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			return conf, err
		}

		viper.ReadConfig(bytes.NewBuffer(content))
	} else {
		viper.AddConfigPath("/etc/gojobs/")
		viper.AddConfigPath("$HOME/.gojobs")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			// load default config
			viper.ReadConfig(bytes.NewBuffer(defaultConf))
		}
	}

	conf.Core.MongoURI = viper.GetString("core.mongo_uri")
	conf.Core.Port = viper.GetString("core.port")
	conf.Core.ZipkinURI = viper.GetString("core.zipkin_uri")
	conf.Core.Storage = viper.GetString("core.storage")

	return conf, nil
}
