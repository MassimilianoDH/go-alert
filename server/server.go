package server

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

var alert models.Alert

func webhook(c *gin.Context) {

	var err error
	var b bytes.Buffer

	if err = c.BindJSON(&alert); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Println("Server: Incoming Alert from: ", alert.Incident.ScopingProjectID)

	data := models.Message{
		ProjectID:    alert.Incident.ScopingProjectID,
		ResourceType: alert.Incident.ResourceTypeDisplayName,
		PolicyName:   alert.Incident.PolicyName,
		ThreatLevel:  alert.Incident.PolicyUserLabels.UserLabel1,
		Summary:      alert.Incident.Summary,
		URL:          alert.Incident.URL,
	}

	t, err := template.ParseFiles("templates/message.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(&b, &data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server: Message sent")
	bot.ExpBot.Send(context.Background(), "TOPIC: Alert", b.String())

	c.JSON(http.StatusOK, alert)
}

func StartServer() {
	r := gin.Default()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		config.Username: config.Password,
	}))

	authorized.POST("/webhook", webhook)

	if err := r.Run(config.Port); err != nil {
		log.Fatal(err)
	}
}
