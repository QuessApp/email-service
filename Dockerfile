# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

RUN go mod init consumer-email-manager

COPY . .

RUN go build -o consumer-email-manager

EXPOSE 8080

CMD [ "./consumer-email-manager" ]