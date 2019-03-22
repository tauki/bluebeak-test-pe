package main

import (
	"fmt"
	"github.com/tauki/bluebeak-test-pe/models"
	"github.com/tauki/bluebeak-test-pe/router"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"time"
)

var cfg *models.Config

func init() {

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
	}

}

func main() {

	// Get the Current Date During Building
	currentTime := time.Now().Local()

	app := cli.NewApp()
	app.Name = "BlueBeak test PE"
	app.Version = fmt.Sprintf("%s-alpha", currentTime.Format("06.01.02")) // get server version by date

	//Cli for Server
	app.Commands = []cli.Command{
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
	}

	app.Run(os.Args)

}

func errorMessage(err error, context string) {
	msg := fmt.Sprintf("ApiGateway :: %s :: %s", context, err.Error())
	fmt.Println(msg)
}

// todo: a better error handling for all packages
