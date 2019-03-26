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
	`USE blue;`,
	`CREATE TABLE IF NOT EXISTS reviews (
		points VARCHAR(255) NOT NULL,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		taster_name VARCHAR(255) NULL,
		taster_twitter_handle VARCHAR(255) NULL,
		price INT NULL,
		designation VARCHAR(255) NULL,
		variety VARCHAR(255) NULL,
		region_1 VARCHAR(255) NULL,
		region_2 VARCHAR(255) NULL,
		province VARCHAR(255) NULL,
		country VARCHAR(45) NULL,
		winery TEXT NOT NULL
	)  ENGINE=InnoDB DEFAULT CHARSET=utf8`,
	`CREATE TABLE IF NOT EXISTS userinfo (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NULL,
		description TEXT NULL,
		profile_image_url VARCHAR(255) NULL,
		followers_count INT UNSIGNED NOT NULL,
		PRIMARY KEY (id)
	)  ENGINE=InnoDB DEFAULT CHARSET=utf8`,
}

var dropTableStatements = []string{
	`DROP TABLE IF EXISTS reviews;`,
	`DROP TABLE IF EXISTS userinfo;`,
}

var dropDBStatements = []string{`DROP DATABASE IF EXISTS blue;`}

type MigrationService struct {
	cfg *models.Config
	db *sql.DB
}

func GetMigrationService(cfg *models.Config) (*MigrationService, error) {

	mysql, err := connection.GetMySqlService(cfg)
	if err != nil {
		msg := fmt.Sprintf("MySQL :: Migration :: Error : %s", err.Error())
		return nil, errors.New(msg)
	}

	return &MigrationService{
		cfg: cfg,
		db:mysql.Conn,
	}, nil
}

// InitMigrate checks database health and confirms database and table schemas
func (ms *MigrationService) InitMigrate() error {
	mysql, err := connection.GetMySqlService(ms.cfg)
	if err != nil {
		msg := fmt.Sprintf("MySQL :: Migration :: Error : %s", err.Error())
		return errors.New(msg)
	}
	defer mysql.Conn.Close()

	// ensure the db is alive
	err = mysql.Conn.Ping()
	if err != nil {
		msg := fmt.Sprintf("MySQL :: Migration :: Error : %s", err.Error())
		return errors.New(msg)
	}

	err = ms.DropDb()
	if err != nil {
		msg := fmt.Sprintf("MySQL :: Migration :: Error : %s", err.Error())
		return errors.New(msg)
	}

	err = ms.ConfirmDatabase()
	if err != nil {
		msg := fmt.Sprintf("MySQL :: Migration :: Error : %s", err.Error())
		return errors.New(msg)
	}

	return err
}

// confirmDatabase creates tables id doesn't exist and if needed the database
func (ms *MigrationService) ConfirmDatabase() error {
	for _, stmt := range createTableStatements {
		_, err := ms.db.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

// dropTables drops the tables in DB
func (ms *MigrationService) DropTables() error {
	for _, stmt := range dropTableStatements {
		_, err := ms.db.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

// dropDb drops database
// caution: it removes all the tables and it's contents
func (ms *MigrationService) DropDb() error {
	for _, stmt := range dropDBStatements {
		_, err := ms.db.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}
