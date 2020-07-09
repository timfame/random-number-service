FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go get -d -v

RUN go build -o main .

FROM alpine

COPY --from=builder /build/main/ /app/

WORKDIR /app 

EXPOSE 8008

ENTRYPOINT ["./main"]
