package tests

import (
	"github.com/tauki/bluebeak-test-pe/connection"
	"github.com/tauki/bluebeak-test-pe/scripts"
	"github.com/tauki/bluebeak-test-pe/services"
	"github.com/tauki/bluebeak-test-pe/services/interfaces"
	"testing"
)

func TestJsonMysqlMigration(t *testing.T) {
	//pwd, _ := os.Getwd()
	cfg := Config

	mysql, err := connection.GetMySqlService(cfg)
	if err != nil {
		t.Error(err)
	}

	reviewCtrl := scripts.GetJsonMysqlMigrationService(cfg, mysql)

	err = reviewCtrl.Execute()
	if err != nil {
		t.Error(err)
	}
}

func TestQueryReviews(t *testing.T) {

	cfg := Config

	t.Log(cfg)

	mysql, _ := connection.GetMySqlService(cfg)

	var dbService interfaces.DbService
	dbService = services.GetDbService(cfg, mysql.Conn)

	rev, _ := dbService.GetReviews(`WHERE NOT taster_name = ""`)

	t.Log(len(rev))

	for _, r := range rev {
		if r.Price != nil {
			t.Log(*r.Price)
		}
	}
}
