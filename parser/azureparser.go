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

func AzureParseAndSend(c *gin.Context) {

	var err error
	var azureAlert models.AzureAlert
	var b bytes.Buffer

	err = c.BindJSON(&azureAlert)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	}

	log.Println("Server: Incoming Alert from: ", azureAlert.SchemaID)

	data := models.AzureMessage{
		AlertID:     azureAlert.Data.Essentials.AlertID,
		SignalType:  azureAlert.Data.Essentials.SignalType,
		AlertRule:   azureAlert.Data.Essentials.AlertRule,
		ThreatLevel: azureAlert.Data.AlertContext.Properties.Name1,
	}

	t, err := template.ParseFiles("templates/azuretemplate.txt")
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
	c.JSON(http.StatusOK, azureAlert)
}
