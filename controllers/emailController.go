package controllers

import (
	dtos "apl-api-email-mocked/dtos"
	serv "apl-api-email-mocked/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostEmail(c *gin.Context) {
	var newEvent dtos.EmailDocumentCanceledDto

	// Call BindJSON to bind the received JSON to
	// newEvent.
	if err := c.BindJSON(&newEvent); err != nil {
		return
	}

	// Add the new event to the slice.
	serv.SendEmailCanceled(newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func ConsumerStarter() {
	serv.SendEmailDocumentReady("br.com.example.signature.document.ready")
}
