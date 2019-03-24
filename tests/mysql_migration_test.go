package tests

import (
	"github.com/tauki/bluebeak-test-pe/scripts"
	"testing"
)

func TestInitMigrate(t *testing.T) {
	migration := scripts.GetMigrationService(Config)

	err := migration.InitMigrate()
	if err != nil {
		t.Error(err)
	}
}
