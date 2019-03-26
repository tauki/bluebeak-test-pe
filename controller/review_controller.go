package controller

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tauki/bluebeak-test-pe/models"
	"github.com/tauki/bluebeak-test-pe/services"
	"github.com/tauki/bluebeak-test-pe/services/interfaces"
	"net/http"
	"strconv"
)

type reviewCtrl struct {
	cfg       *models.Config
	dbService interfaces.DbService
}

func GetReviewController(cfg *models.Config, mysql *sql.DB) *reviewCtrl {

	var dbService interfaces.DbService
	dbService = services.GetDbService(cfg, mysql)

	return &reviewCtrl{
		cfg:       cfg,
		dbService: dbService,
	}
}

func (r *reviewCtrl) GetReviews(c *gin.Context) {

	// todo: fix pagination, doesn't meet requirements

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	var offset int
	if page == 1 {
		offset = 0
	} else {
		offset = limit * page
	}

	query := fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)

	reviews, err := r.dbService.GetReviews(query)
	if err != nil || len(reviews) == 0 {
		r.errorHandler(c, http.StatusNotFound, err.Error())
		return
	}

	next := fmt.Sprintf("%s%s?page=%d", c.Request.Host, "/review", page+1)

	reviewRes := models.DbResponds{
		Data: reviews,
		Next:    next,
	}

	c.JSON(http.StatusFound, reviewRes)
}

func (r *reviewCtrl) AddReview(c *gin.Context) {
	var review models.Reviews

	if err := c.ShouldBindJSON(&review); err != nil {
		r.errorHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	err := r.dbService.InsertReviews(review)
	if err != nil {
		r.errorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, review)
}

func (r *reviewCtrl) errorHandler(router *gin.Context, code int, msg string) {
	router.JSON(code, &models.Message{
		Code:    code,
		Message: msg,
	})
}
