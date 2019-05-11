FROM golang:1.12.5-alpine3.9 as builder

WORKDIR /app
COPY templates main.go go.mod ./
RUN go build -o ./program && chmod +x ./program

FROM golang:1.12.5-alpine3.9
COPY --from=builder /app/templates /app/program /app/

CMD ["/app/program"]
