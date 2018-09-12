package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-squads/genrevan-scheduler/migration"

	"github.com/gorilla/handlers"

	"github.com/urfave/cli"

	"github.com/go-squads/genrevan-scheduler/model"
	"github.com/go-squads/genrevan-scheduler/router"
)

func main() {
	initCLI()
}

func startServer() {
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "PATCH"})

	model.SetupDatabase("development")
	router := router.SetupRouter()
	err := http.ListenAndServe(":8000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router))
	if err != nil {
		fmt.Println(err)
	}
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
				err := migration.RunMigration("development")
				if err == nil {
					fmt.Println("Migration finished")
				} else {
					fmt.Println(err)
				}

				return err
			},
		},
		{
			Name:        "seed",
			Description: "Run database seeder",
			Action: func(c *cli.Context) error {
				err := migration.RunSeeder("development")
				if err == nil {
					fmt.Println("Seeding database finished")
				} else {
					fmt.Println(err)
				}

				return err
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
