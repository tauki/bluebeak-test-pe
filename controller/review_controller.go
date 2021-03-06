package controller

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tauki/bluebeak-test-pe/models"
	"github.com/tauki/bluebeak-test-pe/scripts"
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
		Next: next,
	}

	c.JSON(http.StatusFound, reviewRes)
}

func (r *reviewCtrl) GetUniqueReviewers(c *gin.Context) {

	script, err := scripts.GetMiscService(r.cfg)
	if err != nil {
		r.errorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	reviewers, err := script.UniqueReviewers()
	if err != nil {
		r.errorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusFound, reviewers)

}

func (r *reviewCtrl) GetUserReview(c *gin.Context) {

	query := fmt.Sprintf(`WHERE taster_name = "%s"`, c.Param("name"))

	reviews, err := r.dbService.GetReviews(query)
	if err != nil {
		r.errorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(reviews) == 0 {
		c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	c.JSON(http.StatusFound, reviews)
}

func (r *reviewCtrl) GetUsersWith5ReviewsOrMore(c *gin.Context) {

	script, err := scripts.GetMiscService(r.cfg)
	if err != nil {
		r.errorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	reviews, err := script.UsersWith5ReviewsOrMore()
	if err != nil {
		r.errorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(reviews) == 0 {
		c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	c.JSON(http.StatusFound, reviews)
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

func (r *reviewCtrl) errorHandler(c *gin.Context, code int, msg string) {
	c.JSON(code, &models.Message{
		Code:    code,
		Message: msg,
	})
}
