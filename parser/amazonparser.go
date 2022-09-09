package parser

import (
	"bytes"
	"context"
	"go-alert/bot"
	"go-alert/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AmazonParseAndSend(c *gin.Context) {

	var err error
	var amazonAlert models.AmazonAlert
	var amazonSub models.AmazonSubscriptionConfirmation
	var b bytes.Buffer

	origin := c.Request.Header.Get("x-amz-sns-message-type")

	if origin == "Notification" {
		err = c.BindJSON(&amazonAlert)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		}

		log.Println("Server: Incoming Alert from: ", amazonAlert.TopicArn)

		data := models.AmazonMessage{
			Topic:     amazonAlert.TopicArn,
			Subject:   amazonAlert.Subject,
			Message:   amazonAlert.Message,
			MessageID: amazonAlert.MessageID,
		}

		t, err := template.ParseFiles("templates/amazontemplate.txt")
		if err != nil {
			log.Println(err)
		}

		err = t.Execute(&b, &data)
		if err != nil {
			log.Println(err)
		}

		err = bot.ExpBot.Send(
			context.Background(),
			"🚨🚨🚨  *ALERT*  🚨🚨🚨",
			b.String(),
		)
		if err != nil {
			log.Println(err)
		}

		log.Println("Server: Message sent")
	} else {
		err = c.BindJSON(&amazonSub)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		}

		log.Println("Server: Incoming Subscription Request from: ", amazonSub.TopicArn)
		log.Println("SubscribeURL: ", amazonSub.SubscribeURL)
	}

	c.JSON(http.StatusOK, amazonAlert)
}
