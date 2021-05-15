package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func handleWebhook(c *gin.Context) {
	c.JSON(http.StatusOK, buildFulfillment())
}

func main() {
	var err error
	r := gin.Default()
	r.POST("/webhook", handleWebhook)
	if err = r.Run(); err != nil {
		log.WithError(err).Fatal("Couldn't start server")
	}
}

