# syntax=docker/dockerfile:1
FROM golang:1.18-alpine

WORKDIR /app

COPY . .
RUN go mod download

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -buildvcs=false -o=starwars-api-go

EXPOSE 8080

CMD [ "./starwars-api-go" ]
