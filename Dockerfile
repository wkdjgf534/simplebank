# Build stage
FROM golang:1.23-alpine3.19
WORKDIR /app
COPY . .
RUN go build -o main main.go

EXPOSE 8080
CMD [ "/app/main" ]
