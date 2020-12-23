FROM golang:latest

# ENV 设置环境变量
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE=on

WORKDIR $GOPATH/src/learngo
COPY . $GOPATH/src/learngo
RUN go build .

EXPOSE 8088
ENTRYPOINT ["./learngo"]