package tests

import (
	"github.com/tauki/bluebeak-test-pe/scripts"
	"testing"
)

// todo: configure this test to be run on a dev env

func TestConfirmDatabase(t *testing.T) {

	migration, _ := scripts.GetMigrationService(Config)

	err := migration.ConfirmDatabase()
	if err != nil {
		t.Error(err)
	}
}

func TestDropTables(t *testing.T) {

	migration, _ := scripts.GetMigrationService(Config)

	err := migration.DropTables()
	if err != nil {
		t.Error(err)
	}
}

func TestDropDb(t *testing.T) {

	migration, _ := scripts.GetMigrationService(Config)

	err := migration.DropDb()
	if err != nil {
		t.Error(err)
	}
}

func TestInitMigrate(t *testing.T) {

	migration, _ := scripts.GetMigrationService(Config)

	err := migration.InitMigrate()
	if err != nil {
		t.Error(err)
	}
}
