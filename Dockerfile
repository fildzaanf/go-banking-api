FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go-banking-api ./cmd/api/main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /go-banking-api /go-banking-api

EXPOSE 8080

ENTRYPOINT ["/go-banking-api"]