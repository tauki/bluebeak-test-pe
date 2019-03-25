package main

import (
	"fmt"
	"github.com/tauki/bluebeak-test-pe/models"
	"github.com/tauki/bluebeak-test-pe/router"
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
