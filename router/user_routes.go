package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/tauki/bluebeak-test-pe/controller"
	"github.com/tauki/bluebeak-test-pe/models"
)

func InitUserInfoRouter(router *gin.Engine, cfg *models.Config, mysql *sql.DB) {

	userInfoCtrl := controller.GetUserInfoController(cfg, mysql)

	review := router.Group("user")

	review.POST("", userInfoCtrl.AddReview)
	review.GET("", userInfoCtrl.GetReviews)
}
