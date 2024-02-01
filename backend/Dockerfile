FROM golang:1.21
LABEL authors="zhefu"
WORKDIR /go/src/go-server
COPY . .
RUN go mod tidy
RUN go build -o /go/bin/go-server
EXPOSE 8443
ENTRYPOINT ["/go/bin/go-server"]