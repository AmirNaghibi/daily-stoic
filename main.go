package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//Response is the entire JSON payload response
type Response struct {
	Payload Payload `json:"payload"`
}

//Payload is a google defined higher structure
type Payload struct {
	Google Google `json:"google"`
}

//Google is another google defined higher structure
type Google struct {
	ExpectUserResponse bool         `json:"expectUserResponse"`
	RichResponse       RichResponse `json:"richResponse,omitempty"`
}

//RichResponse gives UI representation of the quote data fetched
type RichResponse struct {
	Items []Item `json:"items,omitempty"`
}

//Item provides different interactions
type Item struct {
	SimpleResponse *SimpleResponse `json:"simpleResponse,omitempty"`
	CarouselBrowse *CarouselBrowse `json:"carouselBrowse,omitempty"`
}

//SimpleResponse provides audio only feedback
type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
}

//CarouselBrowse provides a UI list of items with hyperlinks
type CarouselBrowse struct {
	Items []CarouselItem `json:"items"`
}

//CarouselItem is a UI entry for each stoic quote
type CarouselItem struct {
	Title         string        `json:"title"`
	OpenURLAction OpenURLAction `json:"openUrlAction"`
	Description   string        `json:"description,omitempty"`
}

//OpenURLAction provides a url to be opened in the client's browser when touched
type OpenURLAction struct {
	URL string `json:"url"`
}

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
