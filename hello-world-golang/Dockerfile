FROM  golang:1.14.3  AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

WORKDIR /root

COPY hello-world.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /deploy hello-world.go

FROM alpine:3.7
RUN apk add tzdata ca-certificates && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata && rm -rf /var/cache/apk/* 
COPY --from=builder /deploy /bin/deploy
ENTRYPOINT ["/bin/deploy"]
