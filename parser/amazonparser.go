// NOT WORKING PROPERLY
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
	var b bytes.Buffer

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
	c.JSON(http.StatusOK, amazonAlert)
}
