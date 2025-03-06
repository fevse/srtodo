FROM golang:alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux

COPY . .

RUN go build -o app ./cmd/srtodo/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .
COPY --from=builder /app/db/migrations/20250213210914_new_task_table.sql .
COPY --from=builder /app/.env .

CMD ["./app"]