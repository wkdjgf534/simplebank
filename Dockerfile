# Build stage
FROM golang:1.23-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY start.sh .
COPY wait-for.sh .
COPY app.env .
COPY db/migration ./migration

EXPOSE 8080
# CMD works together with ENTRYPOINT. In this case CMD is an additional params.
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
