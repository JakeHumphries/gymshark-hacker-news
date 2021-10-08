# User
FROM alpine:3.13.1 as user
ARG uid=10001
ARG gid=10001
RUN echo "apiuser:x:${uid}:${gid}::/home/apiuser:/bin/sh" > /apipasswd

# Build
FROM golang:1.16-alpine as build

ENV GO111MODULE=on
WORKDIR /app

COPY . .

RUN target=/go/pkg/mod,sharing=locked \
    && CGO_ENABLED=0 go build -a -o /main cmd/api/main.go

COPY --from=user /apipasswd /etc/passwd
USER apiuser

EXPOSE 8000

CMD [ "/main" ]