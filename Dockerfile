FROM golang:1.9-alpine as builder

WORKDIR /go/src/github.com/Foodee/demo

COPY . /go/src/github.com/Foodee/demo

RUN go build -o demo

CMD ./demo

FROM alpine

COPY --from=builder /go/src/github.com/Foodee/demo/demo /usr/bin/demo

CMD demo
