FROM golang:1.16-alpine

ENV GO111MODULE=on
WORKDIR /app

COPY . .

RUN target=/go/pkg/mod,sharing=locked \
    && CGO_ENABLED=0 go build -a -o /main cmd/consumer/main.go

CMD [ "/main" ]