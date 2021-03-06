package scripts

import (
	"errors"
	"fmt"
	"github.com/tauki/bluebeak-test-pe/connection"
	"github.com/tauki/bluebeak-test-pe/models"
	"github.com/tauki/bluebeak-test-pe/services"
	"github.com/tauki/bluebeak-test-pe/services/interfaces"
)

type Misc struct {
	cfg       *models.Config
	dbService interfaces.DbService
}

// GetMiscServices returns a pointer to the Misc object and error if occurred
// @param: takes in an object of config model
func GetMiscService(cfg *models.Config) (*Misc, error) {
	mysql, err := connection.GetMySqlService(cfg)
	if err != nil {
		msg := fmt.Sprintf("MISC :: Error : %s", err.Error())
		return nil, errors.New(msg)
	}

	var dbService interfaces.DbService
	dbService = services.GetDbService(cfg, mysql.Conn)

	return &Misc{
		dbService: dbService,
		cfg:       cfg,
	}, nil
}

// UsersWith5ReviewsOrMore returns a list of reviewers who has made
// 5 or more reviews
func (m *Misc) UsersWith5ReviewsOrMore() ([]string, error) {
	r, err := m.dbService.QuerySingleCol(
		"taster_name",
		"reviews",
		"GROUP BY taster_name",
		"HAVING COUNT(taster_name) > 1",
		`AND NOT taster_name = ""`,
	)

	if err != nil {
		msg := fmt.Sprintf("MISC :: Error : %s", err.Error())
		return nil, errors.New(msg)
	}

	return r, nil
}

// UniqueReviewers returns a list of unique reviewers (taster_name col)
// from the reviews table
func (m *Misc) UniqueReviewers() ([]string, error) {
	r, err := m.dbService.QuerySingleCol(
		"DISTINCT taster_name",
		"reviews",
		`WHERE NOT taster_name = ""`,
	)

	if err != nil {
		msg := fmt.Sprintf("MISC :: Error : %s", err.Error())
		return nil, errors.New(msg)
	}

	return r, nil
}
