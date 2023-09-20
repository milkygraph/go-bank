FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add curl
RUN go build -o main .
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz --output migrate.tar.gz
RUN tar -xf migrate.tar.gz

FROM alpine:latest AS production
EXPOSE 8080
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 /usr/bin/migrate
COPY db/migration ./migration
COPY wait-for.sh .
COPY start.sh .
COPY app.env .
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
