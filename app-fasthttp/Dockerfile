FROM golang:latest as builder
WORKDIR /go/src/github.com/malagolichiara/benchmarks/fasthttp
COPY main.go .
RUN go mod init malagolichiara/benchmarks && \
    go mod tidy && cat go.mod
RUN go get github.com/valyala/fasthttp
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/github.com/malagolichiara/benchmarks/fasthttp/app .
EXPOSE 8000
CMD ["./app"]

