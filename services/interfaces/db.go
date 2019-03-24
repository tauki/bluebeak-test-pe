package interfaces

import "github.com/tauki/bluebeak-test-pe/models"

// db crud Interface
type DbService interface {
	InsertReviews(review ...models.Reviews) error
	InsertUserInfo(userInfo ...models.UserInfo) error
	GetReviews(args ...string) ([]models.Reviews, error)
}
