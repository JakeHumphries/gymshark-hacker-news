# User
FROM alpine:3.13.1 as user
ARG uid=10001
ARG gid=10001
RUN echo "grpcuser:x:${uid}:${gid}::/home/grpcuser:/bin/sh" > /grpcpasswd

# Build
FROM golang:1.16-alpine as build

ENV GO111MODULE=on
WORKDIR /app

COPY . .

RUN target=/go/pkg/mod,sharing=locked \
    && CGO_ENABLED=0 go build -a -o /main cmd/grpc/main.go

COPY --from=user /grpcpasswd /etc/passwd
USER grpcuser

EXPOSE 9000

CMD [ "/main" ]