package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tauki/bluebeak-test-pe/connection"
	"github.com/tauki/bluebeak-test-pe/models"
	"github.com/tauki/bluebeak-test-pe/router/middleware/cors"
	"net/http"
)

func InitRouter(cfg *models.Config) (*gin.Engine, error) {
	router := gin.New()

	router.Use(cors.CORSMiddleware())
	router.Use(gin.Logger())

	// Setup No Route Message
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Route Not Found"})
	})

	// do an initial mysql health-check
	mySql, err := connection.GetMySqlService(cfg)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s \n", err.Error()))
	}

	err = mySql.Ping()
	if err != nil {
		fmt.Println(fmt.Sprintf("Mysql Service is Offline :  %s \n", err.Error()))
	}

	InitReviewRouter(router, cfg, mySql.Conn)
	InitUserInfoRouter(router, cfg, mySql.Conn)

	return router, nil
}
