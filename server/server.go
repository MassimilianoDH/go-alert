package server

import (
	"go-alert/config"
	"go-alert/parser"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	var err error

	// Create a gin instance
	r := gin.Default()

	// Basic Auth
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		config.Username: config.Password,
	}))

	// Set handles
	if config.GCPTemplate != "" {
		authorized.POST("/gcp", parser.GoogleParseAndSend)
	}

	if config.AZRTemplate != "" {
		authorized.POST("/azr", parser.AzureParseAndSend)
	}

	if config.AWSTemplate != "" {
		authorized.POST("/amazonwebservices", parser.AmazonParseAndSend)
	}

	// listen and serve on 0.0.0.0:config.Port
	err = r.Run(config.Port)
	if err != nil {
		log.Fatal(err)
	}
}
