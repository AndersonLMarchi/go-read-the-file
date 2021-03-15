FROM golang:1.16

WORKDIR /bin
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "read-file.go"]
# CMD ["go", "run", "teste.go"]