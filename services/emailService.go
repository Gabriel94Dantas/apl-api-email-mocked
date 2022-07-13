package services

import (
	contexts "apl-api-email-mocked/contexts"
	dtos "apl-api-email-mocked/dtos"

	"fmt"
)

func SendEmailCanceled(emailEvent dtos.EmailDocumentCanceledDto) {
	fmt.Printf("Email sended for stakeholders from document number : %v\n", emailEvent.IdDocument)
}

func SendEmailDocumentReady(topic string) {
	c := contexts.KafkaConsumer()

	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			emailEventDto := dtos.JsonToEmailEventDto(string(msg.Value))
			var emailDto dtos.EmailDto
			emailDto = dtos.JsonToEmailDto(emailEventDto.Data)
			for i := 0; i < len(emailDto.Emails); i++ {
				fmt.Printf("Email sended to %s document %s is ready.\n", emailDto.Emails[i], emailDto.IdDocument)
			}
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
