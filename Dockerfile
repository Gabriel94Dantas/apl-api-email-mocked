FROM golang:1.18
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o apl-api-email-mocked
RUN rm -rf controllers/
RUN rm -rf services/
RUN rm -rf contexts/
RUN rm -rf models/
RUN rm -rf .vscode/
RUN rm docker-compose.yml
RUN rm Dockerfile
RUN rm main.go
EXPOSE 88
CMD [ "./apl-api-email-mocked" ]