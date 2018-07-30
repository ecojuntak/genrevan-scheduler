package config

import (
	"errors"

	"github.com/go-squads/genrevan-scheduler/util"
	"github.com/spf13/viper"
)

var basepath = util.GetRootFolderPath()

func GetConfig(configFilename string) error {
	viper.SetConfigFile(basepath + "config/" + configFilename)
	if err := viper.ReadInConfig(); err != nil {
		return errors.New("File not found")
	}

	return nil
}
