package parser

import (
	"bytes"
	"context"
	"go-alert/bot"
	"go-alert/config"
	"go-alert/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoogleParseAndSend(c *gin.Context) {

	var err error
	var googleAlert models.GoogleAlert
	var b bytes.Buffer

	err = c.BindJSON(&googleAlert)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	}

	log.Println("Server: Incoming Alert from: ", googleAlert.Incident.ScopingProjectID)

	t, err := template.ParseFiles(config.GCPTemplate)
	if err != nil {
		log.Println(err)
	}

	err = t.Execute(&b, &googleAlert)
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
	c.JSON(http.StatusOK, googleAlert)
}
