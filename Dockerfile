# -------------------
# Builder
# -------------------
FROM golang:1.24-alpine AS builder

# Для локали и tzdata
RUN apk add --no-cache tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Копируем .env чтобы godotenv.Load("../../.env") работал
COPY .env ../../.env

RUN go build -o main ./cmd/app

# -------------------
# Final image
# -------------------
FROM alpine:3.19

# Устанавливаем tzdata для TimeZone
RUN apk add --no-cache tzdata

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/../../.env ../../.env

EXPOSE 8080

CMD ["./main"]
