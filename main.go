package main

import (
	"fmt"
	"github.com/tauki/bluebeak-test-pe/connection"
	"github.com/tauki/bluebeak-test-pe/models"
	"github.com/tauki/bluebeak-test-pe/router"
	"github.com/tauki/bluebeak-test-pe/scripts"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var cfg *models.Config

func init() {
	twiApiKey, twiApiSecret, twiAccessToken, twiAccessTokenSecret := getTwitterKeys()
	cfg = &models.Config{
		ServePort: "9010",
		TLS:       "false",
		TLSPort:   "4443",

		DBName: "blue",
		DBUser: "blue",
		DBPass: "blue",
		DBHost: "127.0.0.1",
		DBPort: "3306",

		CertPrivateKey: "keys/certkey.key",
		CertPath:       "keys/cert.crt",

		JSONPath: "data/winemag-data-130k-v2-formatted.json",

		TwitterAPIKey:            twiApiKey,
		TwitterAPISecret:         twiApiSecret,
		TwitterAccessToken:       twiAccessToken,
		TwitterAccessTokenSecret: twiAccessTokenSecret,
	}

}

func main() {

	// Get the Current Date During Building
	currentTime := time.Now().Local()

	// placeholders for flag
	var script string

	app := cli.NewApp()
	app.Name = "BlueBeak test PE"
	app.Version = fmt.Sprintf("%s-alpha", currentTime.Format("06.01.02")) // get server version by date
	app.Commands = []cli.Command{
		//Cli for Server
		{
			Name:  "run",
			Usage: "Run the server",
			Action: func(c *cli.Context) {

				// Init Router
				route, err := router.InitRouter(cfg)
				if err != nil {
					errorMessage(err, "Router")
				}

				if cfg.TLS == "true" {
					fmt.Println("Serving Gateway on : ", cfg.TLSPort)
					err = http.ListenAndServeTLS(fmt.Sprintf(":%s", cfg.TLSPort), cfg.CertPath, cfg.CertPrivateKey, route)
					if err != nil {
						errorMessage(err, "Serve")
					}
				} else {
					fmt.Println("Serving Gateway on : ", cfg.ServePort)
					err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.ServePort), route)
					if err != nil {
						errorMessage(err, "Serve")
					}
				}
			},
		},
		// scripts
		{
			Name:  "mysql-migrate, msm",
			Usage: "Migrate mysql, create database or tables if doesn't exist",
			Action: func(c *cli.Context) {
				script, err := scripts.GetMigrationService(cfg)
				if err != nil {
					errorMessage(err, "msm")
				}

				err = script.InitMigrate()
				if err != nil {
					errorMessage(err, "msm")
				}
			},
		},
		{
			Name:  "drop-db, dd",
			Usage: "Drop database, deletes all the tables and the stored data",
			Action: func(c *cli.Context) {
				script, err := scripts.GetMigrationService(cfg)
				if err != nil {
					errorMessage(err, "dd")
				}

				err = script.DropDb()
				if err != nil {
					errorMessage(err, "dd")
				}
			},
		},
		{
			Name:  "drop-tables, dt",
			Usage: "Drop tables, deletes all the tables and the stored data",
			Action: func(c *cli.Context) {
				script, err := scripts.GetMigrationService(cfg)
				if err != nil {
					errorMessage(err, "dd")
				}

				err = script.DropTables()
				if err != nil {
					errorMessage(err, "dd")
				}
			},
		},
		{
			Name:  "json-mysql, jm",
			// todo : use flag for sample data location
			Usage: "Read the sample data from the configured path and insert them into db",
			Action: func(c *cli.Context) {
				script, err := scripts.GetJsonMysqlMigrationService(cfg)
				if err != nil {
					errorMessage(err, "jm")
				}

				err = script.Execute()
				if err != nil {
					errorMessage(err, "jm")
				}
			},
		},
		{
			Name:  "users-with-5-reviews-or-more, uw5rm",
			Usage: "Prints out the users in the db with 5 reviews or more",
			Action: func(c *cli.Context) {
				script, err := scripts.GetMiscService(cfg)
				if err != nil {
					errorMessage(err, "uw5rm")
				}

				users, err := script.UsersWith5ReviewsOrMore()
				if err != nil {
					errorMessage(err, "uw5rm")
				}

				for _, user := range users {
					fmt.Println(user)
				}
			},
		},
		{
			Name:  "user-unique, uu",
			Usage: "Prints out the users in the db that are unique",
			Action: func(c *cli.Context) {
				script, err := scripts.GetMiscService(cfg)
				if err != nil {
					errorMessage(err, "uu")
				}

				users, err := script.UniqueReviewers()
				if err != nil {
					errorMessage(err, "uu")
				}

				for _, user := range users {
					fmt.Println(user)
				}
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		errorMessage(err, "CliRun")
	}

}

func errorMessage(err error, context string) {
	msg := fmt.Sprintf("ApiGateway :: %s :: %s", context, err.Error())
	fmt.Println(msg)
}

// getTwitterKeys require a file named twitter-access in the a package directory
// inside the named keys inside the root package directory
// the file should contain, new-line separated and in sequence:
//		API key
//		API secret
//		Access token
//		Access token secret
func getTwitterKeys() (string, string, string, string) {
	f, err := ioutil.ReadFile("keys/twitter-access")
	if err != nil {
		errorMessage(err, "getTwitterKeys")
	}

	creds := strings.Split(string(f), "\n")
	if len(creds) < 4 {
		panic("check credentials")
	}

	fmt.Println(creds)
	// suffix of the strings may contain an extra \n
	return strings.TrimSuffix(creds[0], "\n"),
		strings.TrimSuffix(creds[1], "\n"),
		strings.TrimSuffix(creds[2], "\n"),
		strings.TrimSuffix(creds[3], "\n")
}

// todo: a better error handling for all packages
