FROM golang:1.17 as builder
WORKDIR /go/src/devkit
COPY vendor vendor
COPY go.mod go.mod
COPY go.sum go.sum
COPY controllers controllers
COPY routers routers
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix cgo -ldflags "-w -s -X main.App=devkit" -o app .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /root
COPY --from=builder /go/src/devkit/app .
EXPOSE 8080
VOLUME ["views","static"]

CMD ["./app"]
