FROM golang:1.22.5-alpine
WORKDIR /app

COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080

CMD ["./my-go-app"]

