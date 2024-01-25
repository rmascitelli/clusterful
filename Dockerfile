# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLE=0 go build -o clusterApp .

RUN chmod +x /app/clusterApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/clusterApp /app

ENTRYPOINT [ "/app/clusterApp" ]