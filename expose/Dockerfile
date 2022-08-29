FROM golang:1.18-alpine

RUN mkdir app
COPY main.go app

EXPOSE 8000

CMD go run app/main.go