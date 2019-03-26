package interfaces

import (
	"github.com/tauki/bluebeak-test-pe/models"
)

// db crud Interface
type DbService interface {
	// reviews
	InsertReviews(review ...models.Reviews) error
	GetReviews(args ...string) ([]models.Reviews, error)

	// userInfo
	InsertUserInfo(userInfo ...models.UserInfo) error
	GetUserInfo(args ...string) ([]models.UserInfo, error)

	// misc
	QuerySingleCol(col string, table string, args ...string) ([]string, error)
}
