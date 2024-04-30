ARG  ARCH=amd64
ARG  OS=linux

FROM golang:alpine3.19 AS builder

ENV GOARCH=${ARCH}
ENV GOOS=${OS}

WORKDIR /opt

COPY . .

RUN go mod download
RUN go mod verify
RUN GOARCH=${GOARCH} GOOS=${GOOS} \
    go build -v -o znuny-exporter .

FROM alpine:3.19

COPY --from=builder /opt/znuny-exporter /usr/bin/
COPY --from=builder /opt/config/znuny-exporter.yaml /etc/

EXPOSE 9587

CMD ["znuny-exporter"]
