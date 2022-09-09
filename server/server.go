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
	authorized.POST("/googlecloud", parser.GoogleParseAndSend)
	authorized.POST("/azurecloud", parser.AzureParseAndSend)
	// NOT WORKING PROPERLY
	authorized.POST("/amazonwebservices", parser.AmazonParseAndSend)

	// listen and serve on 0.0.0.0:config.Port
	err = r.Run(config.Port)
	if err != nil {
		log.Fatal(err)
	}
}
