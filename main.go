package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-squads/genrevan-scheduler/migration"

	"github.com/urfave/cli"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/go-squads/genrevan-scheduler/router"
)

func main() {
	initCLI()
}

func startServer() {
	model.SetupDatabase("development")
	router := router.SetupRouter()
	http.ListenAndServe(":8000", router)
}

func initCLI() {
	cliApp := cli.NewApp()
	cliApp.Name = "Genrevan Scheduler App"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:        "migrate",
			Description: "Run database migration",
			Action: func(c *cli.Context) error {
				fmt.Println("Migration finished")
				return migration.RunMigration("development")
			},
		},
		{
			Name:        "seed",
			Description: "Run database seeder",
			Action: func(c *cli.Context) error {
				fmt.Println("Seeding finished")
				return migration.RunSeeder("development")
			},
		},
		{
			Name:        "start",
			Description: "Start REST API Server",
			Action: func(c *cli.Context) error {
				fmt.Println("Genreven Scheduler started at port 8000...")
				startServer()
				return nil
			},
		},
	}

	cliApp.Run(os.Args)
}
