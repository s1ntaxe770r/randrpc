FROM golang:alpine as builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev
COPY . .

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux  go build -o client 

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/client  /client 

EXPOSE 7070

ENTRYPOINT ["/client"]

