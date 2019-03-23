package scripts

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tauki/bluebeak-test-pe/connection"
	"github.com/tauki/bluebeak-test-pe/models"
	"github.com/tauki/bluebeak-test-pe/services"
	"github.com/tauki/bluebeak-test-pe/services/interfaces"
	"io/ioutil"
	"os"
)

type JsonMysqlMigration struct {
	cfg       *models.Config
	dbService interfaces.DbService
}

func GetJsonMysqlMigrationService(cfg *models.Config, mysql *connection.MySqlService) *JsonMysqlMigration {

	var dbService *services.DbService
	dbService = services.GetDbService(cfg, mysql.Conn)

	return &JsonMysqlMigration{
		cfg:       cfg,
		dbService: dbService,
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

	//fmt.Println(len(reviews))

	//for _, i := range reviews {
	//	fmt.Printf("%+v\n", i)
	//}

	// todo: error handle on partial failure

	// prepare batch
	batch := make([]models.Reviews, 0)

	for i, review := range reviews {
		batch = append(batch, review)
		if len(batch) == 2000 || i == len(reviews)-1 {

			err := j.dbService.InsertReviews(batch...)
			if err != nil {
				msg := fmt.Sprintf("Reviews :: insertion :: %s", err.Error())
				return errors.New(msg)
			}

			batch = nil
		}
	}

	return nil
}
