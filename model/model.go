package model

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-squads/genrevan-scheduler/config"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func SetupDatabase(env string) error {
	var connection string
	var conf config.Conf
	if env == "development" {
		c, err := conf.GetConfig(env + ".yaml")
		if err != nil {
			return err
		}

		connection = fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=%s", c.DbUser, c.DbPassword, c.DbName, c.DbSslMode)
	} else if env == "testing" {
		c, err := conf.GetConfig(env + ".yaml")
		if err != nil {
			return err
		}

		connection = fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=%s", c.DbUser, c.DbPassword, c.DbName, c.DbSslMode)
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
