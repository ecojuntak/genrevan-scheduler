package config

import (
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Conf struct {
	DbHost     string `yaml:"DB_HOST"`
	DbName     string `yaml:"DB_NAME"`
	DbUser     string `yaml:"DB_USER"`
	DbPassword string `yaml:"DB_PASSWORD"`
	DbSslMode  string `yaml:"DB_SSL_MODE"`
}

func (c *Conf) GetConfig(configFilename string) (*Conf, error) {
	yamlFile, err := ioutil.ReadFile("../config/" + configFilename)
	if err != nil {
		return nil, errors.New("File not found")
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
