# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLE=0 go build -o rpcApp .

RUN chmod +x /app/rpcApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/rpcApp /app

CMD [ "/app/rpcApp" ]