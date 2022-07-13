package dtos

import (
	"encoding/json"
)

type JSON map[string]interface{}

type EmailEventDto struct {
	Id              string `json:"id"`
	SpecVersion     string `json:"specVersion"`
	Source          string `json:"source"`
	Type            string `json:"type"`
	Subject         string `json:"subject"`
	Time            string `json:"time"`
	CorrelationID   string `json:"correlationId"`
	DataContentType string `json:"dataContentType"`
	Data            JSON   `json:"data"`
}

func JsonToEmailEventDto(jsonString string) EmailEventDto {
	data := EmailEventDto{}
	json.Unmarshal([]byte(jsonString), &data)
	return data
}
