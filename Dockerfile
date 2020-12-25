FROM scratch

# ENV 设置环境变量
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE=on

WORKDIR $GOPATH/src/learngo
COPY . $GOPATH/src/learngo

EXPOSE 8089
CMD ["./learngo"]