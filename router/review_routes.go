package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/tauki/bluebeak-test-pe/controller"
	"github.com/tauki/bluebeak-test-pe/models"
)

func InitReviewRouter(router *gin.Engine, cfg *models.Config, mysql *sql.DB) {

	reviewCtrl := controller.GetReviewController(cfg, mysql)

	review := router.Group("review")

	review.POST("", reviewCtrl.AddReview)
	review.GET("", reviewCtrl.GetReviews)
	review.GET("unique", reviewCtrl.GetUniqueReviewers)
}
