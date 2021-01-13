FROM golang:alpine as builder 
RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . . 

RUN go mod download
RUN CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o shippy-user-cli *.go 

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app 
WORKDIR /app

COPY --from=builder /app/shippy-user-cli .

ENTRYPOINT [ "./shippy-user-cli" ]