# User
FROM alpine:3.13.1 as user
ARG uid=10001
ARG gid=10001
RUN echo "consumeruser:x:${uid}:${gid}::/home/consumeruser:/bin/sh" > /consumerpasswd

FROM golang:1.16-alpine as build

ENV GO111MODULE=on
WORKDIR /app

COPY . .

RUN target=/go/pkg/mod,sharing=locked \
    && CGO_ENABLED=0 go build -a -o /main cmd/consumer/main.go

COPY --from=user /consumerpasswd /etc/passwd
USER consumeruser

CMD [ "/main" ]