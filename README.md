# apl-api-email-mocked
## Introduction

This application has the responsability of consumer events to notice all stakeholders that a document is ready. The document is ready because evereyone has signed. Another functinality we have here is to send a email when a document is canceled. 

## Features

- An endpoint that send a email to stakeholders when a document is canceled.
- A event consumer that notice all stakeholders a document is ready.

## Tech

apl-api-email-mocked was wrote in Golang 1.18 and we use this dependencies:

```
require (
	github.com/confluentinc/confluent-kafka-go v1.9.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.8.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20211007075335-d3039528d8ac // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
```

- gin-goinc is a lib that allowed us to receive requests
- we use confluent lib client for go to connect our application to Kafka broker

## Installation

First of all you have to clone this repository and inside the correct folder you have to do this command:

```
go run main.go
```

## Development

This project was initiated by Gabriel Araujo Dantas a Brazilian Computer Engineer and my colleague Lucas Henrique A. de Paula a Brazilian Computer Scientist.

## Docker

If you want to run all features we are using you can use the docker-compose-all-in-one.yml, and you have to follow these steps:

- When you work with two docker-compose you need to create the network first (
  ```
  docker network create --driver=bridge  --subnet=172.18.0.0/16  --ip-range=172.18.0.0/24  --gateway=172.18.0.1   my_network
  ```
  ), this is necessary because you have use same network on all docker-compose
  
- After that you have to add on your kafka docker-compose this (
```
networks: 
  default: 
    external: 
      name: kafka_confluent_network
```
      )
- And now you have to run all docker-compose and you are ready to use this API
```
docker-compose up -d
```
or
```
docker-compose up -d --build
```
.

## Important information

This application is a part of events plataform, so you have to create all other parts of this plataform

Another important information is we use the confluent kafka community for this test. You can find that on this link: https://github.com/confluentinc/cp-all-in-one/tree/7.0.1-post/cp-all-in-one-community 

We use as interface REST to kafka our other project:
https://github.com/Gabriel94Dantas/api-eventos
