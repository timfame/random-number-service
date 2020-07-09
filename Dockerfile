FROM golang

WORKDIR $GOPATH/src/github.com/timfame/random_number_service

COPY . .

RUN go get -v -u ./...

RUN go build -o main .

EXPOSE 8008

ENTRYPOINT ["./main"]
