package dtos

import "encoding/json"

type EmailDto struct {
	Emails     []string
	IdDocument string
}

func JsonToEmailDto(jsonMap map[string]interface{}) EmailDto {
	var emailDto EmailDto
	jsonByte, err := json.Marshal(jsonMap)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonByte, &emailDto)
	return emailDto
}
