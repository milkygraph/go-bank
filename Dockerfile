FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 8080

FROM alpine:latest AS production
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/app.env .
CMD [ "/app/main" ]
