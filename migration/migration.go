package migration

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/go-squads/genrevan-scheduler/util"
)

var basepath = util.GetRootFolderPath()

func RunMigration(env string) error {
	err := model.SetupDatabase(env)
	if err != nil {
		return err
	}

	migrationQueryString, err := GetStringFromFile(basepath + "migration/schema.sql")

	if err != nil {
		return err
	}

	_, err = model.Db.Exec(*migrationQueryString)

	if err != nil {
		return err
	}

	return nil
}

func RunSeeder(env string) error {
	err := model.SetupDatabase(env)
	if err != nil {
		return err
	}

	seederQueryString, err := GetStringFromFile(basepath + "migration/seeder.sql")

	if err != nil {
		return err
	}

	model.Db.Exec(*seederQueryString)

	return nil
}

func GetStringFromFile(filename string) (*string, error) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, errors.New("File not found")
	}

	queryString := fmt.Sprintf("%s", content)

	return &queryString, nil
}
