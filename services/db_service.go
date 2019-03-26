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

	insert, err := tx.Prepare(insertReviewsStatement)
	if err != nil {
		msg := fmt.Sprintf("dbService :: txPrepare :: %s", err.Error())
		return errors.New(msg)
	}

	for _, review := range reviews {
		_, err := insert.Exec(
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

func (db *DbService) GetReviews(args ...string) ([]models.Reviews, error) {

	query := fmt.Sprintf("SELECT * FROM reviews \n")

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

const insertUserInfoStatement = `
  INSERT INTO userinfo (
    name, description, profile_image_url, followers_count
  ) VALUES (?, ?, ?, ?)`

func (db *DbService) InsertUserInfo(userInfos ...models.UserInfo) error {
	//todo twitter API access is required
	// twitter get follower count https://cdn.syndication.twimg.com/widgets/followbutton/info.json?screen_names=

	tx, err := db.conn.Begin()
	if err != nil {
		msg := fmt.Sprintf("dbService :: tx :: %s", err.Error())
		return errors.New(msg)
	}
	defer tx.Rollback()

	insert, err := tx.Prepare(insertUserInfoStatement)
	if err != nil {
		msg := fmt.Sprintf("dbService :: txPrepare :: %s", err.Error())
		return errors.New(msg)
	}

	for _, userInfo := range userInfos {
		_, err := insert.Exec(
			&userInfo.Id,
			&userInfo.Name,
			&userInfo.Description,
			&userInfo.ProfileImageUrl,
			&userInfo.FollowerCount)
		if err != nil {
			msg := fmt.Sprintf("dbService :: InsertUserInfo :: %s", err.Error())
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

func (db *DbService) GetUserInfo(args ...string) ([]models.UserInfo, error) {

	query := fmt.Sprintf("SELECT * FROM userinfo \n")

	for _, arg := range args {
		query += arg + "\n"
	}

	rows, err := db.conn.Query(query)
	if err != nil {
		msg := fmt.Sprintf("dbService :: GetUserInfo :: %s", err.Error())
		return nil, errors.New(msg)
	}

	var userInfos []models.UserInfo
	for rows.Next() {
		var userInfo models.UserInfo
		if err := rows.Scan(
			&userInfo.Id,
			&userInfo.Name,
			&userInfo.Description,
			&userInfo.ProfileImageUrl,
			&userInfo.FollowerCount); err != nil {
			return nil, err
		}

		userInfos = append(userInfos, userInfo)
	}

	return userInfos, nil
}

func (db *DbService) QuerySingleCol(col string, table string, args ...string) ([]string, error) {

	query := fmt.Sprintf("SELECT %s FROM %s", col, table)

	for _, arg := range args {
		query += arg + "\n"
	}

	rows, err := db.conn.Query(query)
	if err != nil {
		msg := fmt.Sprintf("dbService :: QuerySingleCol :: %s", err.Error())
		return nil, errors.New(msg)
	}

	var data []string

	for rows.Next() {
		var datum string
		if err := rows.Scan(&datum); err != nil {
			return nil, err
		}

		data = append(data, datum)
	}

	return data, nil
}

// todo : db-query with go type assessment
// see go reflection laws
