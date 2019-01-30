FROM golang:alpine3.8 AS builder
WORKDIR /go/src/github.com/wzshiming/openapi/
COPY . .
RUN go install github.com/wzshiming/openapi/cmd/...

FROM wzshiming/upx AS upx
COPY --from=builder /go/bin/ /go/bin/
RUN upx /go/bin/*

FROM alpine:3.8
RUN apk add -U --no-cache ca-certificates openssl tzdata
COPY --from=upx /go/bin/ /usr/local/bin/

EXPOSE 9000

CMD openapi