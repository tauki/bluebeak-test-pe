package tests

import "github.com/tauki/bluebeak-test-pe/models"

var Config *models.Config

func init() {
	Config = &models.Config{

		DBName: "blue",
		DBUser: "blue",
		DBPass: "blue",
		DBHost: "127.0.0.1",
		DBPort: "3306",

		CertPrivateKey: "keys/certkey.key",
		CertPath:       "keys/cert.crt",

		JSONPath: "../data/winemag-data-130k-v2-formatted.json",
	}
}