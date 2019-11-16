
FROM golang:1.13 as builder
ADD apis /supportbot/apis
ADD go.mod /supportbot/go.mod
ADD cmd /supportbot/cmd
ADD pkg /supportbot/pkg
ADD main.go /supportbot/main.go
ADD vendor /supportbot/vendor
WORKDIR /supportbot/
RUN go build -o  supportbot main.go

FROM ubuntu:18.04
COPY --from=builder  /supportbot/supportbot /supportbot
ENTRYPOINT ["/supportbot"]
