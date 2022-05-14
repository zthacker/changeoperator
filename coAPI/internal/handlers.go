package coAPI

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func (co *ChangeOperator) Init() {
	go co.HandleRequests()
}

func (co *ChangeOperator) HandleRequests() {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/change/create", co.AddChange)

	router.Run(":" + os.Getenv("CHANGEOPERATOR_PORT"))
}
