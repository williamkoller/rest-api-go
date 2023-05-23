FROM golang:1.20-alpine
LABEL authors="william"

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy
