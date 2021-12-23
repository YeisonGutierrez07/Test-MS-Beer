FROM golang:latest AS builder

RUN apt-get update && apt-get install -y ca-certificates openssl

ARG cert_location=/usr/local/share/ca-certificates

# Get certificate from "github.com"
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org"
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Update certificates
RUN update-ca-certificates


RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src
COPY go.mod .
RUN go mod download
COPY . .
RUN go build main.go
EXPOSE 5000

FROM scratch
COPY --from=builder /go/src .
ENTRYPOINT  ["./main"]


# // docker build -t yeisongutty/test_ms_beer:1.0.0 .
# // docker run -d -p 5000 yeisongutty/test_ms_beer:1.0.0