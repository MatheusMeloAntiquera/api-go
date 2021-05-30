FROM golang:latest

WORKDIR /go/api-go
EXPOSE 8080

# RUN go get github.com/canthefason/go-watcher
# RUN go install github.com/canthefason/go-watcher/cmd/watcher@latest

RUN go get github.com/pilu/fresh

ENTRYPOINT [ "docker/golang/entrypoint.sh" ]