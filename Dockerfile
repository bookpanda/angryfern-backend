FROM golang:1.21.4-alpine3.18 as builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go


FROM alpine AS runner
WORKDIR /app

COPY --from=builder /app/server ./

ENV GO_ENV production

EXPOSE 3000

CMD ["./server"]