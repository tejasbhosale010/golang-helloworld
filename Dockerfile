FROM golang:1.22.5
WORKDIR /app

COPY . .

RUN go build -o my-go-app .

CMD ["./my-go-app"]

