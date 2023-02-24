# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

RUN go mod init email-service

COPY . .

RUN go build -o email-service

EXPOSE 8080

CMD [ "./email-service" ]