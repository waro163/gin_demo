FROM golang:1.16

ENV TZ Asia/Shanghai
ENV GOPROXY "https://goproxy.cn,direct"

COPY . /workdir
WORKDIR /workdir

RUN go mod vendor
RUN go build -mod=vendor -o main main.go

EXPOSE 8080
ENTRYPOINT ["./main"]