package config

import (
	"errors"
	"io/ioutil"

	"github.com/go-squads/genrevan-scheduler/util"
	yaml "gopkg.in/yaml.v2"
)

type Conf struct {
	DbHost     string `yaml:"DB_HOST"`
	DbName     string `yaml:"DB_NAME"`
	DbUser     string `yaml:"DB_USER"`
	DbPassword string `yaml:"DB_PASSWORD"`
	DbSslMode  string `yaml:"DB_SSL_MODE"`
}

var basepath = util.GetRootFolderPath()

func (c *Conf) GetConfig(configFilename string) (*Conf, error) {
	yamlFile, err := ioutil.ReadFile(basepath + "/config/" + configFilename)
	if err != nil {
		return nil, errors.New("File not found")
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
