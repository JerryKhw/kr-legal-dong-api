FROM golang:latest AS build
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main main.go

FROM alpine:latest
WORKDIR /
COPY --from=build /app/main ./main
COPY --from=build /app/kr-legal-dong ./kr-legal-dong
RUN chmod +x main

EXPOSE 8080

ENTRYPOINT [ "./main" ]
