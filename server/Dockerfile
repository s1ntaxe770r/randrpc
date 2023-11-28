FROM golang:1.18 as builder

WORKDIR /app

COPY server/go.mod .
COPY proto . 
RUN go mod download
#RUN apk add --no-cache gcc musl-dev build-base

COPY . .

WORKDIR /app/server
RUN GOOS=linux GOARCH=amd64 go build  -o api 

## FROM alpine


## COPY --from=builder /app/api /api

EXPOSE 7070


ENTRYPOINT ["/app/server/api"]

