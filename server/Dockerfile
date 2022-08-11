
FROM golang:alpine 

RUN apk add --update git

ADD . $GOPATH/src/build/

WORKDIR $GOPATH/src/build/

RUN go mod download

EXPOSE 8081

CMD ["go", "run", "main.go"]