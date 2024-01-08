FROM golang:alpine AS build
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev
COPY . .
RUN CGO_ENABLED=1 CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go build -o main main.go

FROM alpine:latest
WORKDIR /
COPY --from=build /app/main ./main
COPY --from=build /app/kr-legal-dong ./kr-legal-dong
RUN chmod +x main

EXPOSE 8080

ENTRYPOINT [ "./main" ]
