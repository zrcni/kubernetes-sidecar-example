FROM golang:1.12.5-alpine3.9 as builder

WORKDIR /app
COPY templates ./templates
COPY main.go go.mod ./
RUN go build -o ./program && chmod +x ./program

FROM golang:1.12.5-alpine3.9
COPY --from=builder /app/program /app/
COPY --from=builder /app/templates /app/templates

CMD ["/app/program"]
