FROM golang:1.16

WORKDIR /go/src/api-go
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...


EXPOSE 8080