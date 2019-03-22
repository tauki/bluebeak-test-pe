package services

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/tauki/bluebeak-test-pe/models"
)

type DbService struct {
	// until db is decoupled
	conn *sql.DB
	cfg  *models.Config
}

func GetDbService(cfg *models.Config, conn *sql.DB) *DbService {

	return &DbService{
		conn: conn,
		cfg:  cfg,
	}
}

const insertReviewsStatement = `
  INSERT INTO reviews (
    points, title, description, taster_name, taster_twitter_handle, price, designation, variety, region_1, region_2, province, country, winery
  ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

func (db *DbService) InsertReviews(reviews *[]models.Reviews) error {
	tx, err := db.conn.Begin()
	if err != nil {
		msg := fmt.Sprintf("dbService :: tx :: %s", err.Error())
		return errors.New(msg)
	}
	defer tx.Rollback()

	insertReviews, err := tx.Prepare(insertReviewsStatement)
	if err != nil {
		msg := fmt.Sprintf("dbService :: txPrepare :: %s", err.Error())
		return errors.New(msg)
	}

	for _, review := range *reviews {

		_, err := insertReviews.Exec(
			review.Points,
			review.Title,
			review.Description,
			review.TasterName,
			review.TasterTwitterHandle,
			review.Price,
			review.Designation,
			review.Variety,
			review.Region1,
			review.Region2,
			review.Province,
			review.Country,
			review.Winery,
		)
		if err != nil {
			msg := fmt.Sprintf("dbService :: InsertReviewsBatch :: %s", err.Error())
			return errors.New(msg)
		}
	}

	err = tx.Commit()
	if err != nil {
		msg := fmt.Sprintf("dbService :: txCommit :: %s", err.Error())
		return errors.New(msg)
	}

	return nil
}

const insertUserInfoStatement = `
  INSERT INTO userinfo (
    name, description, profile_image_url, followers_count
  ) VALUES (?, ?, ?, ?)`

func (db *DbService) InsertUserInfo(userInfo *[]models.UserInfo) error {
	//todo
	return nil
}
