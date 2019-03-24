package services

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/tauki/bluebeak-test-pe/models"
)

//todo: pool, pipeline

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

func (db *DbService) InsertReviews(reviews ...models.Reviews) error {
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

	for _, review := range reviews {
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
			msg := fmt.Sprintf("dbService :: InsertReviews :: %s", err.Error())
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

func (db *DbService) InsertUserInfo(userInfo ...models.UserInfo) error {
	//todo
	// twitter get follower count https://cdn.syndication.twimg.com/widgets/followbutton/info.json?screen_names=
	return nil
}

func (db *DbService) QueryReviews(col string, args ...string) ([]models.Reviews, error) {

	query := fmt.Sprintf("SELECT %s FROM reviews \n", col)

	for _, arg := range args {
		query += arg + "\n"
	}

	rows, err := db.conn.Query(query)
	if err != nil {
		msg := fmt.Sprintf("dbService :: QueryDb :: %s", err.Error())
		return nil, errors.New(msg)
	}

	var reviews []models.Reviews
	for rows.Next() {
		var review models.Reviews
		if err := rows.Scan(
			&review.Points,
			&review.Title,
			&review.Description,
			&review.TasterName,
			&review.TasterTwitterHandle,
			&review.Price,
			&review.Description,
			&review.Variety,
			&review.Region1,
			&review.Region2,
			&review.Province,
			&review.Country,
			&review.Winery); err != nil {
			return nil, err
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}
