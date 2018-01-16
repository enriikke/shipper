FROM golang:1.9 as builder

WORKDIR /go/src/github.com/enriikke/shipper/consignment-cli
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/enriikke/shipper/consignment-cli/consignment-cli .

CMD ["./consignment-cli"]
