FROM golang:latest

WORKDIR /go/src/reader
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["reader"]