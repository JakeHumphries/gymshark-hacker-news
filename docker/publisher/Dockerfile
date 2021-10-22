# User
FROM alpine:3.13.1 as user
ARG uid=10001
ARG gid=10001
RUN echo "publisheruser:x:${uid}:${gid}::/home/publisheruser:/bin/sh" > /publisherpasswd

FROM golang:1.16-alpine as build

ENV GO111MODULE=on
WORKDIR /app

COPY . .

RUN target=/go/pkg/mod,sharing=locked \
    && CGO_ENABLED=0 go build -a -o /main cmd/publisher/main.go

COPY --from=user /publisherpasswd /etc/passwd
USER publisheruser

CMD [ "/main" ]