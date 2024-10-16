FROM golang:1.23-alpine AS golang

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN rm -rf tmp
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tmp/main .

FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=golang /app/tmp/main main

EXPOSE 8080

CMD ["./main"]
