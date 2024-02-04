FROM golang:1.21-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependencies
COPY ["app/go.mod", "./"]
RUN go mod download

# build
COPY app ./
RUN go build -o ./bin/app cmd/app.go

FROM alpine as runnner

COPY --from=builder /usr/local/src/bin/app /
COPY app/.env /.env

EXPOSE 10000

CMD ["/app"]