package services

import (
	"database/sql"
	"github.com/tauki/bluebeak-test-pe/models"
)

type DbService struct {
	// until db is decoupled
	conn *sql.DB
	cfg *models.Config
}

func GetDbService(cfg *models.Config ,conn *sql.DB) *DbService {
	return &DbService{
		conn: conn,
		cfg: cfg,
	}
}

func (db *DbService) InsertReviews(review *models.Reviews) error {
	return nil
}

func (db *DbService) InsertUserInfo(userInfo *models.UserInfo) error {
	return nil
}
