package tests

import (
	"github.com/tauki/bluebeak-test-pe/connection"
	"testing"
)

func TestMysqlService(t *testing.T) {

	cfg := *Config

	mysql, err := connection.GetMySqlService(&cfg)
	if err != nil {
		t.Error(err)
	}

	if err = mysql.Ping(); err != nil {
		t.Fatal("The Mysql server is not alive")
	}
	if err = mysql.Conn.Close(); err != nil {
		t.Error(err)
	}

	cfg.DBUser = "random"
	mysql, err = connection.GetMySqlService(&cfg)
	if err == nil {
		t.Errorf("Expected the function to fail for Username: %s", cfg.DBUser)
	}

	cfg = *Config
	cfg.DBName = "random"
	mysql, err = connection.GetMySqlService(&cfg)
	if err == nil {
		t.Errorf("Expected the function to fail for DBName: %s", cfg.DBName)
	}

	cfg = *Config
	cfg.DBPort = "random"
	mysql, err = connection.GetMySqlService(&cfg)
	if err == nil {
		t.Errorf("Expected the function to fail for Port: %s", cfg.DBPort)
	}

	cfg = *Config
	cfg.DBPass = "random"
	mysql, err = connection.GetMySqlService(&cfg)
	if err == nil {
		t.Errorf("Expected the function to fail for Password: %s", cfg.DBPass)
	}

	cfg = *Config
	cfg.DBHost = "random"
	mysql, err = connection.GetMySqlService(&cfg)
	if err == nil {
		t.Errorf("Expected the function to fail for Host: %s, you got the wrong database", cfg.DBHost)
	}
}
