FROM golang:alpine as builder

WORKDIR /app
COPY templates ./templates
COPY main.go go.mod ./

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -o ./program \
  && chmod +x ./program \
  && apk add -U --no-cache ca-certificates

FROM scratch

COPY --from=builder /app/program /program
COPY --from=builder /app/templates /templates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/program"]
