package coAPI

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (co *ChangeOperator) AddChange(g *gin.Context) {
	var attrs Attrs
	validate := validator.New()

	if err := g.BindJSON(&attrs); err != nil {
		log.Error(err)
		g.JSON(http.StatusBadRequest, gin.H{"error binding json": err.Error()})
		return
	}

	validationErr := validate.Struct(attrs)
	if validationErr != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error validating struct": validationErr.Error()})
		return
	}

	result, err := co.PostgresClient.Exec("INSERT INTO changes (requester, env, type, customerimpact, description, date, link, linkback, time) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		attrs.Requester, attrs.Env, attrs.Type, attrs.CustomerImpact, attrs.Description, attrs.Date, attrs.Link, attrs.LinkBack, time.Now())
	if err != nil {
		log.Error(err)
		g.JSON(http.StatusInternalServerError, gin.H{"couldn't insert to postgres": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"successfully created a change": result})
	log.Infof("Change created! %s", result)
}
