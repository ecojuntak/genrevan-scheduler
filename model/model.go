package model

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-squads/genrevan-scheduler/config"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var Db *sql.DB

func SetupDatabase(env string) error {
	var connection string

	if env == "development" {
		err := config.GetConfig(env + ".yaml")
		if err != nil {
			return err
		}

		connection = fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=%s", viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_NAME"), viper.GetString("DB_SSL_MODE"))
	} else if env == "testing" {
		err := config.GetConfig(env + ".yaml")
		if err != nil {
			return err
		}

		connection = fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=%s", viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_NAME"), viper.GetString("DB_SSL_MODE"))
	} else {
		return errors.New("Environment not match")
	}

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	Db = db

	return nil
}
