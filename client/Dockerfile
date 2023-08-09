FROM golang:alpine as builder

WORKDIR /app
COPY ./go.mod .
RUN go mod download
RUN apk add --no-cache gcc musl-dev
COPY . .
RUN CGO_ENABLED=1 GOOS=linux  go build -o client

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/client /client

EXPOSE 8080

ENTRYPOINT ["/client"]
