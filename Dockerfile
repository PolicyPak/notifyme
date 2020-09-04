FROM golang:1.14-alpine AS builder
WORKDIR /
RUN apk add gcc g++ --no-cache
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"'  notifier.go

FROM scratch
COPY --from=builder /notifier /notifier
ENTRYPOINT ["/notifier"]