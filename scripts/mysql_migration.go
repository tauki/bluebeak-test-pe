package scripts

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/tauki/bluebeak-test-pe/connection"
	"github.com/tauki/bluebeak-test-pe/models"
)

var createTableStatements = []string{
	`CREATE DATABASE IF NOT EXISTS blue DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE library;`,
	`CREATE TABLE IF NOT EXISTS reviews (
		points VARCHAR(255) NOT NULL,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		taster_name VARCHAR(45) NULL,
		taster_twitter_handle VARCHAR(255) NULL,
		price INT NULL,
		designation VARCHAR(255) NULL,
		variety VARCHAR(255) NULL,
		region_1 VARCHAR(45) NULL,
		region_2 VARCHAR(45) NULL,
		province VARCHAR(45) NULL,
		country VARCHAR(45) NULL,
		winery VARCHAR(45) NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS userinfo (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NULL,
		description TEXT NULL,
		profile_image_url VARCHAR(255) NULL,
		followers_count INT UNSIGNED NOT NULL,
		PRIMARY KEY (id)
	)`,
}

type MigrationService struct {
	cfg *models.Config
}

func GetMigrationService(cfg *models.Config) *MigrationService {
	return &MigrationService{
		cfg: cfg,
	}
}

// InitMigrate checks database health and confirms database and table schemas
func (ms *MigrationService) InitMigrate(cfg *models.Config) error {
	mysql, err := connection.GetMySqlService(cfg)
	if err != nil {
		msg := fmt.Sprintf("MySQL :: Migration :: Error : %s", err.Error())
		return errors.New(msg)
	}

	// ensure the db is alive
	err = mysql.Conn.Ping()
	if err != nil {
		msg := fmt.Sprintf("MySQL :: Migration :: Error : %s", err.Error())
		return errors.New(msg)
	}

	return confirmDatabase(mysql.Conn)
}

// confirmDatabase creates tables id doesn't exist and if needed the database
func confirmDatabase(conn *sql.DB) error {
	for _, stmt := range createTableStatements {
		_, err := conn.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}