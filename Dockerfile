FROM golang:latest as builder

WORKDIR /cmd

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /cmd/main .

ENV PORT=:8000
ENV LOG_LEVEL=debug

CMD ["./main"]