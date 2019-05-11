FROM golang:1.12.5-alpine3.9 as builder

WORKDIR /app
COPY main.go go.mod ./
RUN go build -o ./program && chmod +x ./program

FROM golang:1.12.5-alpine3.9
COPY --from=builder /app/program /app/program

CMD ["/app/program"]
