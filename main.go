package main

import (
	control "apl-api-email-mocked/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	go control.ConsumerStarter()
	router := gin.Default()
	router.POST("/email/document/canceled", control.PostEmail)
	router.Run("0.0.0.0:88")
}
