FROM golang:1.12.5-alpine3.9 as builder

WORKDIR /app
COPY main.go go.mod ./

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -o ./program \
  && chmod +x ./program \
  && apk add -U --no-cache ca-certificates

FROM scratch

COPY --from=builder /app/program /program
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/program"]
