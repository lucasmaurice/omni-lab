FROM golang:1.26-alpine as builder

WORKDIR /go/src

COPY go.mod ./go.mod
COPY go.sum ./go.sum
COPY main.go ./main.go

RUN go mod download
RUN go build -o ftpd -ldflags "-s -w" main.go

FROM scratch
LABEL maintainer="Sonic <sonic@djls.io>"
LABEL org.opencontainers.image.source=https://github.com/justereseau/wip
LABEL org.opencontainers.image.description="wip"
LABEL org.opencontainers.image.licenses=WTFPL

COPY --from=builder /go/src/ftpd /bin/ftpd

ENTRYPOINT [ "/bin/ftpd" ]
