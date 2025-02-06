FROM golang:alpine AS builder
ENV GO111MODULE=on
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd/building
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app/main .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder app/. .
WORKDIR /app/cmd/building
COPY --from=builder app/cmd/building .
EXPOSE 8081
CMD ["app/main"]