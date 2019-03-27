package models

type Config struct {
	// Server Information
	TLS       string
	TLSPort   string
	ServePort string // Server Listening Port

	// DB

	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string

	// SSL Credential

	CertPrivateKey string
	CertPath       string

	// JSON data
	JSONPath string

	// Twitter
	Twitter                  string
	TwitterAPIKey            string
	TwitterAPISecret         string
	TwitterAccessToken       string
	TwitterAccessTokenSecret string
	TwitterKeyPath           string
}
