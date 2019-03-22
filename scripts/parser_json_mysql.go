package scripts

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tauki/bluebeak-test-pe/connection"
	"github.com/tauki/bluebeak-test-pe/models"
	"io/ioutil"
	"os"
)

type JsonMysqlMigration struct {
	cfg *models.Config
}

func GetJsonMysqlMigrationService(cfg *models.Config) *JsonMysqlMigration {

	return &JsonMysqlMigration{
		cfg: cfg,
	}
}

func (j *JsonMysqlMigration) Execute() error {

	// open json file
	jsonFile, err := os.Open(j.cfg.JSONPath)
	if err != nil {
		msg := fmt.Sprintf("Reviews :: %s", err.Error())
		return errors.New(msg)
	}
	defer jsonFile.Close()

	// Read the file
	jsonByte, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		msg := fmt.Sprintf("Reviews :: %s", err.Error())
		return errors.New(msg)
	}

	var reviews []models.Reviews
	err = json.Unmarshal(jsonByte, &reviews)
	if err != nil {
		msg := fmt.Sprintf("Reviews :: %s", err.Error())
		return errors.New(msg)
	}

	//for _, i := range reviews {
	//	fmt.Printf("%+v\n", i)
	//}

	// get MySQL controller
	db, err := connection.GetMySqlService(j.cfg)
	if err != nil {
		msg := fmt.Sprintf("Reviews :: %s", err.Error())
		return errors.New(msg)
	}
	defer db.Conn.Close()

	return nil
}
