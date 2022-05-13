# FROM golang:1.16
FROM golang:alpine

ENV GOPROXY "https://goproxy.cn,direct"

COPY . /workdir
WORKDIR /workdir

RUN go mod vendor
RUN go build -mod=vendor -o main main.go

From alpine

ENV TZ Asia/Shanghai
COPY --from=0 /workdir/main /workdir/
COPY ./config.yaml /workdir/
WORKDIR /workdir
RUN apk add bash

EXPOSE 8080
ENTRYPOINT ["./main"]