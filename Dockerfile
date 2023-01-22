FROM golang:1.19-alpine as builder

WORKDIR /app

RUN apk add --update --no-cache ca-certificates

COPY ./go.mod ./main.go ./

RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o docker-entrypoint .

FROM scratch

COPY --from=builder /etc/ssl/certs /etc/ssl/certs

WORKDIR /

COPY --from=builder /app/docker-entrypoint .

ENTRYPOINT [ "./docker-entrypoint" ]
