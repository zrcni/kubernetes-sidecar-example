FROM golang:1.12.5-alpine3.9

WORKDIR /app

CMD ["go", "run", "main.go"]
