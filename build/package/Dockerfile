FROM golang:1.17 as builder

WORKDIR /go/src/app

COPY . .

RUN go build -o /build/bin/api .

FROM centos:7

EXPOSE 8000

COPY --from=builder /build/bin/api /build/bin/api

COPY /upload /build/bin/upload

ENTRYPOINT ["/build/bin/api"]