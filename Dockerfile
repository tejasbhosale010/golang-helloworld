FROM golang:1.22.5-alpine
WORKDIR /app

COPY . .


CMD ["./my-go-app"]

