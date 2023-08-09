FROM golang:alpine as builder

WORKDIR /app
COPY ./go.mod .
RUN go mod download
RUN apk add --no-cache gcc musl-dev
COPY . .

COPY ../proto/  .

RUN old_string="replace github.com/s1ntaxe770r/randrpc/proto => ../proto" 
RUN new_string="replace github.com/s1ntaxe770r/randrpc/proto => ./proto"


RUN sed -i "s|$old_string|$new_string|g"



RUN CGO_ENABLED=1 GOOS=linux  go build -o client 

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/client  /client 

EXPOSE 7070

ENTRYPOINT ["/client"]

