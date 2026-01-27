FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN ls -R /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main /app/cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY dev/.env .

CMD ["./main"]
