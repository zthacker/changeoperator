package coAPI

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func ReadGinJSONBody(g gin.Context) {
	jd, err := ioutil.ReadAll(g.Request.Body)
	if err != nil {
		log.Error(err)
	}
	log.Info(string(jd))
}
